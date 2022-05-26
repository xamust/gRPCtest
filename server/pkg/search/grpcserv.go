package search

import (
	"golang.org/x/net/context"
	"server/pkg/api"
)

type GRPCServer struct{}

func (g *GRPCServer) Search(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{
		Writer: "TestWriter",
		Book:   "TestBook",
	}, nil
}
