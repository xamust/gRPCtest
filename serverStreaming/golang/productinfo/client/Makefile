.PHONY: compile

run:
		make build
		bin/server

build:
		go build -i -v -o bin/server

## Compile the proto file.
compile:
		protoc -I ecommerce/proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. ecommerce/proto/*.proto

.DEFAULT_GOAL := build