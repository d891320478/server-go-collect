package main

import (
	"net/http"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"github.com/d891320478/server-go-collect/bili-danmu/httpmethod"
)

func main() {
	// 初始化日志
	baselog.InitLog("bili-danmu")
	// 初始化rpc
	config.SetConsumerService(bean.BiliRpcService)
	if err := config.Load(); err != nil {
		panic(err)
	}
	http.HandleFunc("/bili-danmu/danmuList", httpmethod.DanmuList)
	http.HandleFunc("/bili-danmu/danmuList/testGift", httpmethod.DanmuListTestGift)
	err := http.ListenAndServe(":9755", nil)
	if err != nil {
		panic(err)
	}
}
