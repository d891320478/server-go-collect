package main

import "github.com/d891320478/server-go-collect/startBiliHttp/bililive"

// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/startBiliHttp -x main.go
func main() {
	bililive.StartBiliHttp()
}
