BINARY_NAME=http-server
.DEFAULT_GOAL := dev

build:
		GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows ./src/main.go
		GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux ./src/main.go
		GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin ./src/main.go

run: build
		./target/${BINARY_NAME}-windows

build_and_run: build run

dev:
		nodemon -x "go run ./src/main.go"
