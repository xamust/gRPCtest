.PHONY: init build run compile

all:
		make init
		make run
init:
		go mod tidy
build:
		make init
		go build -o build/client -v ./cmd/app
run:
		go run ./cmd/app
## Compile the proto file.
compile:
		protoc -I api/proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. api/proto/*.proto

.DEFAULT_GOAL := build