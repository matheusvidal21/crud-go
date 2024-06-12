package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/matheusvidal21/crud-go/src/configuration/logger"
	"github.com/matheusvidal21/crud-go/src/controller"
	"github.com/matheusvidal21/crud-go/src/controller/routes"
	service "github.com/matheusvidal21/crud-go/src/model/service"
	"log"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserController(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
