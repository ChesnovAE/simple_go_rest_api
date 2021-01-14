package main

import (
	"log"
	"os"

	"github.com/ChesnovAE/simple_go_rest_api/internal/server"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/handler"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/repository"
	"github.com/ChesnovAE/simple_go_rest_api/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error occured while initializing config:\n %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while loading env variables:\n %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		DBEngine: viper.GetString("db.dbengine"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("error while app tried connect to db:\n %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server:\n %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("./config")

	return viper.ReadInConfig()
}
