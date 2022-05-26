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

	//инициализируем пустую структуру для ответа клиенту...
	searchResp := &api.SearchResponse{}

	//определяем, что ищет клиент (книги или писателя) и выдаем требуемую информацию, либо возвращаем ошибку...
	if req.GetWriter() != "" {
		//поиск книг...
		book, err := g.Store.FindByWriter(req.GetWriter())
		if err != nil {
			g.Logger.Errorf(err.Error())
			return nil, err
		}
		searchResp.Book = book
	} else if req.GetBook() != "" {
		//поиск писателя...
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
