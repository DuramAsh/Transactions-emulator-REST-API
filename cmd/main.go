package main

import (
	emulator "github.com/duramash/constanta-emulator-task"
	"github.com/duramash/constanta-emulator-task/pkg/handler"
	"github.com/duramash/constanta-emulator-task/pkg/repository"
	"github.com/duramash/constanta-emulator-task/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Error with config initialization: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error with loading env data: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetInt("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize DB: %s", err.Error())
	}

	repo := repository.NewRepository(db)
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
