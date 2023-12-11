APP_NAME = getAllDanmu

buildapp:
	GOOS=linux GOARCH=amd64 go build -o build/${APP_NAME} main.go
	sh ~/script/aliyun/207-scp-aliyun.sh build/${APP_NAME}

push:
	git add .
	git commit -m "c"
	git push