.PHONY: init build run compile test

all:
		make init
		make run
init:
		go mod tidy
build:
		make init
		go build -o build/server -v ./cmd/app
run:
		go run ./cmd/app
## Compile the proto file.
compile:
		protoc -I api/proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. api/proto/*.proto
test:
		go test -v ./...
.DEFAULT_GOAL := build