APP_NAME = bili-danmu

build-app:
	GOOS=linux GOARCH=amd64 go build -o build/${APP_NAME} ./cmd
	sh ~/script/aliyun/207-scp-aliyun.sh build/${APP_NAME}