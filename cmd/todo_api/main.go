package main

import (
	"log"

	"github.com/ChesnovAE/simple_go_rest_api/internal/server"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/handler"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/repository"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error iccured while initializing config %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("./config")

	return viper.ReadInConfig()
}
