package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"server/internal/app/store"
	"server/pkg/api"
	"server/pkg/search"
)

type AppServer struct {
	config  *Config
	logger  *logrus.Logger
	StoreBD *store.AppStore
	server  *grpc.Server
	search  *search.GRPCServer
}

//init new app
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

	//init db...
	if err := s.configureStore(); err != nil {
		return err
	}

	//init grpc...
	if err := s.configureGRPC(); err != nil {
		return err
	}

	return nil
}

//configure mysql store conn...
func (s *AppServer) configureStore() error {
	newStore := store.New(s.config.Store)
	if err := newStore.Open(); err != nil {
		return err
	}
	s.StoreBD = newStore
	s.logger.Info("Подключение к БД установлено успешно!")
	return nil
}

//configure logrus...
func (s *AppServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	s.logger.Info("Инициализация логера выполнена успешно!")
	return nil
}

//configure gRPC...
func (s *AppServer) configureGRPC() error {
	s.server = grpc.NewServer()
	s.search = &search.GRPCServer{s.logger, s.StoreBD.StoreRep()}
	api.RegisterSearchingServer(s.server, s.search)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	if err := s.server.Serve(l); err != nil {
		return err
	}
	s.logger.Info("Запуск gRPC сервера...")
	return nil
}
