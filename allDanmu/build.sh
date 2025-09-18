#!/bin/bash

set -e

APP_NAME=getAllDanmu
GOOS=linux GOARCH=amd64 go build -o build/${APP_NAME} cmd/app.go
cd build
rm -f getAllDanmu.zip
zip getAllDanmu.zip getAllDanmu
if [[ -f ~/script/aliyun/207-scp-aliyun.sh ]];then
    sh ~/script/aliyun/207-scp-aliyun.sh ${APP_NAME}
fi
