go-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/startBiliHttp -x main.go
scp:
	sh ~/script/aliyun/207-scp-aliyun.sh build/startBiliHttp
push:
	git add .
	git commit -m "c"
	git push