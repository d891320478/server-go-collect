package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/d891320478/server-go-collect/allDanmu/bean"
	"github.com/d891320478/server-go-collect/allDanmu/bililive"
	baselog "github.com/d891320478/server-go-collect/base-log"
)

func main() {
	baselog.InitLog("allDanmu")
	config.SetConsumerService(bean.BiliRpcService)
	if err := config.Load(); err != nil {
		panic(err)
	}
	bililive.AllDanMu(222272)
}
