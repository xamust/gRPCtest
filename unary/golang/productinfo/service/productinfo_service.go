package main

import (
	"context"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "productinfo/service/ecommerce"
)

type server struct {
	productMap map[string]*pb.Product
}

func (s server) AddProduct(ctx context.Context, product *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	product.Id = out.String()

	s.productMap[product.Id] = product

	return &pb.ProductID{Value: product.Id}, status.New(codes.OK, "").Err()
}

func (s server) GetProduct(ctx context.Context, id *pb.ProductID) (*pb.Product, error) {
	value, ok := s.productMap[id.Value]
	if ok {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product %v does not exist.", id.Value)
}
