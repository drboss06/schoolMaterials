package main

import (
	"log"
	"schoolMaterial/internal/db"
	"schoolMaterial/internal/repository"
	"schoolMaterial/internal/server"
	service "schoolMaterial/internal/services"
	handler "schoolMaterial/internal/transport/rest"
)

func main() {
	database, err := db.Connect(db.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	repo := repository.NewPostgresMaterialRepository(database)
	materialService := service.NewMaterialService(repo)
	handler := handler.NewHandler(materialService)

	srv := new(server.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}

}
