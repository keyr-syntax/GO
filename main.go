package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keyr-syntax/server/config"
	"github.com/keyr-syntax/server/models"
	"github.com/keyr-syntax/server/routes"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Todo{})
	router := gin.Default()
	routes.RegisterTodoRoutes(router)
	router.Run(":8080")
}