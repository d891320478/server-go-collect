package httpmethod

import (
	"context"
	"net/http"
	"strconv"

	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"github.com/d891320478/server-go-collect/bili-danmu/biliservice"
	"github.com/d891320478/server-go-collect/bili-danmu/domain"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var wu = &websocket.Upgrader{
	ReadBufferSize:  8192,
	WriteBufferSize: 8192,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const dmLen = 30

func DanmuList(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		baselog.ErrorLog().Error("get ws error. err = %v", err)
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	defer ws.Close()
	// 判断roomId白名单
	roomIdStr := r.URL.Query().Get("roomId")
	roomId, err := strconv.ParseInt(roomIdStr, 10, 0)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	roomIdWrapper := &wrapperspb.Int64Value{
		Value: roomId,
	}
	can, err := bean.BiliRpcService.RoomCanUseServer(context.Background(), roomIdWrapper)
	if err != nil || !can.Value {
		baselog.ErrorLog().Error("get ws error. err = %v, roomId = %d, can = %t", err, roomId, can.Value)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// 连接弹幕
	danmu := make(chan domain.DanMuVO)
	closeFlag := make(chan bool)
	defer close(danmu)
	defer close(closeFlag)
	client, err := biliservice.Register(danmu, roomId)
	if err != nil {
		w.WriteHeader(502)
		w.Write([]byte("连接弹幕失败"))
		return
	}
	defer biliservice.CloseClient(client)
	go func() {
		for {
			messageType, _, err := ws.ReadMessage()
			if err != nil || messageType == websocket.CloseMessage {
				closeFlag <- true
				return
			}
		}
	}()
	// 写消息
	var msgList []domain.DanMuVO
	for i := 1; i <= dmLen; i++ {
		msgList = append(msgList, domain.DanMuVO{Empty: true})
	}
	for flag := false; !flag; {
		select {
		case dm := <-danmu:
			if len(dm.Avatar) == 0 {
				dm.Avatar = biliservice.GetAvatar(dm.Uid)
			}
			msgList = append(msgList, dm)
			if len(msgList) > dmLen {
				msgList = msgList[1:]
			}
			err = ws.WriteJSON(msgList)
			if err != nil {
				baselog.ErrorLog().Error("ws write error. err is %v", err)
				flag = true
			}
		case flag = <-closeFlag:
		}
	}
}
