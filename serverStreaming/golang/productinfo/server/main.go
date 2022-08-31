package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "productinfo/server/ecommerce"
)

const port = ":50051"

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{make(map[string]*pb.Order)})
	log.Printf("Starting gRPC server on port, %v\n", port)
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}

}
