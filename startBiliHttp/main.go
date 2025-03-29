package main

import (
	baselog "github.com/d891320478/server-go-collect/base-log"
	"github.com/d891320478/server-go-collect/startBiliHttp/bililive"
)

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/startBiliHttp -x main.go
func main() {
	baselog.InitLog("startBiliHttp")
	bililive.StartBiliHttp()
}
