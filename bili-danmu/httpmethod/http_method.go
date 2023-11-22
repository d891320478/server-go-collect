package httpmethod

import (
	"context"
	"net/http"
	"strconv"
	"time"

	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
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
		return
	}
	defer ws.Close()
	// 判断roomId白名单
	roomIdStr := r.URL.Query().Get("roomId")
	roomId, _ := strconv.ParseInt(roomIdStr, 10, 0)
	roomIdWrapper := &wrapperspb.Int64Value{
		Value: roomId,
	}
	can, err := bean.BiliRpcService.RoomCanUseServer(context.Background(), roomIdWrapper)
	if err != nil || !can.Value {
		return
	}
	// 连接弹幕
	// 写消息
	ws.WriteMessage(websocket.TextMessage, []byte("1"))
	time.Sleep(1 * time.Second)
	ws.WriteMessage(websocket.TextMessage, []byte("2"))
	time.Sleep(1 * time.Second)
	ws.WriteMessage(websocket.TextMessage, []byte("3"))
	time.Sleep(1 * time.Second)
}
