package main

import (
	"log"

	"github.com/hachibu/flipcoin/internal/config"
	"github.com/hachibu/flipcoin/internal/web"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server := web.NewServer(config)
	server.ListenAndServe()
}
