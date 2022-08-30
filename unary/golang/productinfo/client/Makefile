.PHONY: compile

run:
		make build
		bin/client

build:
		go build -i -v -o bin/client

## Compile the proto file.
compile:
		protoc -I ecommerce/proto --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. ecommerce/proto/*.proto

#.DEFAULT_GOAL := build