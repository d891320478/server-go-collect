package main

import (
	"context"

	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/d891320478/server-go-collect/allDanmu/bean"
	"github.com/d891320478/server-go-collect/allDanmu/bililive"
	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/proto-go/bili"
)

func main() {
	baselog.InitLog("allDanmu")

	config.SetConsumerService(bean.BiliRpcService)
	if err := config.Load(); err != nil {
		panic(err)
	}

	req := bili.AddNewGuardRequest{
		RoomId:     int64(222272),
		Uid:        int64(325170),
		Username:   "幽诗人",
		GuardLevel: int32(3),
	}
	_, err0 := bean.BiliRpcService.AddNewGuard(context.Background(), &req)
	if err0 != nil {
		baselog.ErrorLog().Error("AddNewGuard error: %s", err0.Error())
	}

	bililive.AllDanMu(222272)
}
