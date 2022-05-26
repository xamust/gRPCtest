package server

import "server/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",           //default params(если не будет данных параметров в конфиге)...
		LogLevel: "info",            //default params(если не будет данных параметров в конфиге)...
		Store:    store.NewConfig(), //new store config data (конфиг структура для бд)...
	}
}
