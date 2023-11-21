package main

import (
	"net/http"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/d891320478/server-go-collect/base-log/baselog"
	"github.com/d891320478/server-go-collect/bili-danmu/bean"
	"github.com/d891320478/server-go-collect/bili-danmu/httpmethod"
	"github.com/gorilla/mux"
)

func main() {
	// 初始化日志
	baselog.InitLog("bili-danmu")
	// 初始化rpc
	config.SetConsumerService(bean.BiliRpcService)
	if err := config.Load(); err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/bili-danmu/danmuList", httpmethod.DanmuList)
	http.ListenAndServe(":9755", router)

	select {}
}
