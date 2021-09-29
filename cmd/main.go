package main

import (
	"log"

	"github.com/Yoture/link_keeper/core"
	config "github.com/Yoture/link_keeper/internal/api/handlers"
)

func main() {
	handlers := new(config.HttpHandler)
	srv := new(core.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
