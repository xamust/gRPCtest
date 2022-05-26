package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"server/internal/app/server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server.toml", "Path to config file")
}

func main() {

	//парсим configs/server.toml и заправляем данные структуры Config...
	flag.Parse()
	config := server.NewConfig()
	meta, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	//чекаем все ли парметры из конфига распарсились...
	if len(meta.Undecoded()) != 0 {
		log.Fatal("Undecoded configs param: ", meta.Undecoded())
	}

	//стартуем само приложение с вышеполученным конфигом...
	serverGRPC := server.New(config)
	if err := serverGRPC.Start(); err != nil {
		log.Fatal(err)
	}
}
