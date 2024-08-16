package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"schoolMaterial/internal/db"
	"schoolMaterial/internal/repository"
	"schoolMaterial/internal/server"
	service "schoolMaterial/internal/services"
	handler "schoolMaterial/internal/transport/rest"
	"schoolMaterial/pkg/logger"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = logger.InitLogger("app.log", viper.GetString("logLevel"))
	if err != nil {
		log.Fatal("Logger init error", err)
	}

	err = initConfig()
	if err != nil {
		logger.GetLogger().Error(err)
	}

	database, err := db.Connect(db.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.GetLogger().Error(err)
	}
	logger.GetLogger().Info("Database connected")

	defer database.Close()

	repo := repository.NewPostgresMaterialRepository(database)
	materialService := service.NewMaterialService(repo)
	handler := handler.NewHandler(materialService)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logger.GetLogger().Fatal(err)
	}

	logger.GetLogger().Info("Server started")

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
