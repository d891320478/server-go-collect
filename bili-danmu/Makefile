APP_NAME = bili-danmu

build-app:
	GOOS=linux GOARCH=amd64 go build -o build/${APP_NAME} ./cmd

clean:
	rm -rf ./build/${APP_NAME}

tidy:
	go mod tidy

test:
	go test ./...
