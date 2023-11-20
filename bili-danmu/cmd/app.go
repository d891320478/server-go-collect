package main

import (
	"context"
	"fmt"

	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	roomId := &wrapperspb.Int64Value{
		Value: 222272,
	}
	can, err := bean.BiliRpcService.RoomCanUseServer(context.Background(), roomId)
	fmt.Println(can.Value)
	fmt.Println(err)
}
