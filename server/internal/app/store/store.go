package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AppStore struct {
	config    *Config
	db        *sql.DB
	storeRepo *StoreRepo
}

func New(config *Config) *AppStore {
	return &AppStore{
		config: config,
	}
}

func (s *AppStore) Open() error {
	db, err := sql.Open("mysql", s.config.DataBaseUrl)
	if err != nil {
		return err
	}

	//todo: from habr...
	//https://habr.com/ru/company/oleg-bunin/blog/583558/
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil { //ping connect DB
		return err
	}

	s.db = db

	return nil
}

func (s *AppStore) Close() {
	s.db.Close()
}

func (s *AppStore) StoreRep() *StoreRepo {
	if s.storeRepo != nil {
		return s.storeRepo
	}

	s.storeRepo = &StoreRepo{
		store: s,
	}

	return s.storeRepo
}
