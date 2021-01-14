package main

import (
	"log"

	"github.com/ChesnovAE/simple_go_rest_api/internal/server"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server %s", err.Error())
	}
}
