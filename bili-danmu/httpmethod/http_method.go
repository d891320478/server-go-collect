package httpmethod

import (
	"context"
	"net/http"
	"strconv"

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
		return
	}
	defer ws.Close()

	roomIdStr := r.URL.Query().Get("roomId")
	roomId, _ := strconv.ParseInt(roomIdStr, 10, 0)
	roomIdWrapper := &wrapperspb.Int64Value{
		Value: roomId,
	}
	can, err := bean.BiliRpcService.RoomCanUseServer(context.Background(), roomIdWrapper)
	if err != nil || !can.Value {
		return
	}
}
