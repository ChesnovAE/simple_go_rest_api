package main

import (
	"log"

	"github.com/ChesnovAE/simple_go_rest_api/internal/server"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/handler"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/repository"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server %s", err.Error())
	}
}
