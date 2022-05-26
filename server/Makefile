.PHONY: init build run test

all:
		make init
#		make build
		make test
		make run
init:
		go mod tidy
build:
		make init
#		go build -v ./cmd/app
		go build -o build/server -v ./cmd/app
run:
		go run ./cmd/app

#test:
#		go test -v ./...

#protoc -I api/proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. api/proto/*.proto
.DEFAULT_GOAL := build