package httpmethod

import (
	"context"
	"net/http"
	"strconv"

	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"github.com/d891320478/server-go-collect/bili-danmu/biliservice"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

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
	danmu := make(chan string)
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
	ws.SetCloseHandler(func(code int, text string) error {
		closeFlag <- true
		return nil
	})
	go func() {
		for {
			if _, _, err := ws.NextReader(); err != nil {
				closeFlag <- true
				break
			}
		}
	}()
	// 写消息
	for flag := false; !flag; {
		select {
		case val := <-danmu:
			err = ws.WriteMessage(websocket.TextMessage, []byte(val))
			if err != nil {
				baselog.ErrorLog().Error("ws write error. err is %v", err)
				flag = true
			}
		case flag = <-closeFlag:
			baselog.InfoLog().Info("closeFlag = %t", flag)
		}
		baselog.InfoLog().Info("flag = %t", flag)
	}
	baselog.InfoLog().Info("close ws success")
}
