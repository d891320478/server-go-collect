package main

import (
	"context"
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	config.SetConsumerService(bean.BiliRpcService)
	if err := config.Load(); err != nil {
		panic(err)
	}

	roomId := &wrapperspb.Int64Value{
		Value: 222272,
	}
	can, err := bean.BiliRpcService.RoomCanUseServer(context.Background(), roomId)
	fmt.Println(can.Value)
	fmt.Println(err)
}
