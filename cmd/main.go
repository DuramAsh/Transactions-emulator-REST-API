package main

import (
	emulator "github.com/duramash/constanta-emulator-task"
	"github.com/duramash/constanta-emulator-task/pkg/handler"
	"github.com/duramash/constanta-emulator-task/pkg/repository"
	"github.com/duramash/constanta-emulator-task/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error with config initialization: %s", err.Error())
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	server := new(emulator.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
