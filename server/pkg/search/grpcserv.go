package search

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"server/internal/app/store"
	"server/pkg/api"
)

type GRPCServer struct {
	Logger *logrus.Logger
	Store  *store.StoreRepo
}

func (g *GRPCServer) Search(ctx context.Context, req *api.SearchRequest) (*api.SearchResponse, error) {

	searchResp := &api.SearchResponse{}

	if req.GetWriter() != "" {
		book, err := g.Store.FindByWriter(req.GetWriter())
		if err != nil {
			g.Logger.Errorf(err.Error())
			return nil, err
		}
		searchResp.Book = book
	} else if req.GetBook() != "" {

		writer, err := g.Store.FindByBook(req.GetBook())
		if err != nil {
			g.Logger.Errorf(err.Error())
			return nil, err
		}

		searchResp.Writer = writer
	} else {
		g.Logger.Errorf("Параметры book и writer не указаны!")
		return nil, fmt.Errorf("Параметры book и writer не указаны!")
	}

	return searchResp, nil
}
