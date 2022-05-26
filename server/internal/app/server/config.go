package server

import "server/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",           //default params...
		LogLevel: "info",            //default params...
		Store:    store.NewConfig(), //new store config data...
	}
}
