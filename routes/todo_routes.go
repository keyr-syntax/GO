package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/keyr-syntax/server/controllers"
)

func RegisterTodoRoutes(router *gin.Engine) {
	todoController := &controllers.TodoControllers{}
	router.POST("/create",todoController.NewTodo)
	router.GET("/todos", todoController.GetallTodos)
	router.GET("/:id", todoController.FindTodoByID)
	router.PUT("/update/:id", todoController.UpdateTodo)
	router.DELETE("/delete/:id",todoController.DeleteTodo)
}