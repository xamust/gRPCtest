package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"server/internal/app/store"
	"server/pkg/api"
	"server/pkg/search"
)

type AppServer struct {
	config  *Config
	logger  *logrus.Logger
	storeBD *store.AppStore
}

//init new server
func New(config *Config) *AppServer {
	return &AppServer{
		config: config,
		logger: logrus.New(), //sirupsen/logrus
	}
}

func (s *AppServer) Start() error {

	//init log...
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Инициализация логера выполнена успешно!")

	//init db...
	if err := s.configureStore(); err != nil {
		return err
	}
	s.logger.Info("Подключение к БД установлено успешно!")

	sg := grpc.NewServer()
	srv := search.GRPCServer{}
	api.RegisterSearchingServer(sg, &srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	sg.Serve(l)

	return nil
}

func (s *AppServer) configureStore() error {
	newStore := store.New(s.config.Store)
	if err := newStore.Open(); err != nil {
		return err
	}
	s.storeBD = newStore

	return nil
}

//configure logrus...
func (s *AppServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
