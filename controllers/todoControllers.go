package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/keyr-syntax/server/config"
	"github.com/keyr-syntax/server/models"
)

type TodoControllers struct{}

func (tc *TodoControllers) NewTodo(c *gin.Context){
	var todo models.Todo
	
	if err := c.ShouldBindJSON(&todo); err!=nil{
		c.JSON(400, gin.H{"success":false, "message":"Invalid JSON data"})
		return
	}
	if todo.Title == "" || todo.Description == ""{
		c.JSON(400, gin.H{"success":false, "message":"All fields are requred!"})
		return
	}
	if err := config.DB.Create(&todo).Error; err !=nil{
		c.JSON(403, gin.H{"success":false, "message":"Failed to create todo"})
		return
	}else{
		c.JSON(201, gin.H{"success":true, "message":"Todo added","todo":todo})
		return
	}
}

func (tc *TodoControllers) GetallTodos(c *gin.Context){
	var todos []models.Todo
	if err := config.DB.Find(&todos).Error; err != nil{
		c.JSON(404, gin.H{"success":"false", "message":"Todos not found"})
		return
	}
	c.JSON(200, gin.H{"success":"true","todos":todos})
}

func(tc *TodoControllers)FindTodoByID(c *gin.Context){
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err != nil{
		c.JSON(404, gin.H{"success":"false", "message":"Todo not found"})
		return
	}
	c.JSON(200, gin.H{"success":"true","todo":todo})
}

func(uc *TodoControllers)UpdateTodo(c *gin.Context){
	var todo models.Todo
	var updatedTodo models.Todo
	id := c.Param("id")

	if err := c.ShouldBindJSON(&updatedTodo); err !=nil{
		c.JSON(400, gin.H{"success":"false", "message":"Invalid JSON"})
		return
	}
	if updatedTodo.Title == "" || updatedTodo.Description == ""{
		c.JSON(400, gin.H{"success":"false", "message":"All fields are required"})
		return
	}

	if err := config.DB.First(&todo, id).Error; err !=nil{
		c.JSON(404, gin.H{"success":"false","message":"Todo not found"})
		return
	}
	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	config.DB.Save(&todo)
	c.JSON(200, gin.H{"success":"true","message":"Todo updated","Todo":todo})
}


func (uc *TodoControllers)DeleteTodo(c *gin.Context){
	var todo models.Todo
	id := c.Param("id")

	if err := config.DB.First(&todo, id).Error; err !=nil{
		c.JSON(404, gin.H{"success":false, "message":"Todo not found"})
		return
	}
	if err := config.DB.Delete(&todo).Error; err !=nil{
		c.JSON(400,"Failed to delete the Todo")
	} else {
		c.JSON(200, gin.H{"success": true, "message": "Deleted"})
	}


}