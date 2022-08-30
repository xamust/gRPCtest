package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "productinfo/client/ecommerce"
	"time"
)

const (
	address = "localhost"
	port    = "50051"
)

func main() {
	conn, err := grpc.Dial(net.JoinHostPort(address, port), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "TestName1"
	descr := "This text is test info, for testname1"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{
		Name:        name,
		Description: descr,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Product ID: %v added successfully\n", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Product: %s\n", product.String())
}
