package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type AppStore struct {
	config    *Config
	db        *sql.DB
	storeRepo *StoreRepo
}

//инициализируем новый конфиг для коннекта к бд...
func New(config *Config) *AppStore {
	return &AppStore{
		config: config,
	}
}

//открываем бд...
func (s *AppStore) Open() error {
	db, err := sql.Open("mysql", s.config.DataBaseUrl)
	if err != nil {
		return err
	}

	//пингуем бд, для проверки соединения, открытое соединение не гарантирует работоспособность соединения с бд!...
	if err := db.Ping(); err != nil { //ping connect DB
		return err
	}

	s.db = db

	return nil
}

//реализация закрытия соединения с бд...
func (s *AppStore) Close() {
	s.db.Close()
}

//реализация хранилища запросов к бд, для работы сервиса...
func (s *AppStore) StoreRep() *StoreRepo {
	if s.storeRepo != nil {
		return s.storeRepo
	}

	s.storeRepo = &StoreRepo{
		store: s,
	}

	return s.storeRepo
}
