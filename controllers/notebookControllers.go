package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/keyr-syntax/server/config"
	"github.com/keyr-syntax/server/models"
	"gorm.io/gorm"
)

type notebookControllers struct{}

func (nc *notebookControllers) newNotebook(c *gin.Context){
	var notebook models.Notebook
	if err := c.ShouldBindBodyWithJSON(&notebook); err != nil{
		c.JSON(400, gin.H{"success":false, "message":"Invalid JSON"})
		return
	}
	if notebook.TopicID == 0 || notebook.Content == "" || !notebook.IsDraft || !notebook.IsPublished || notebook.Order == 0 || notebook.Title == "" {
		c.JSON(403, gin.H{"success":false, "message":"All fields are required"})
		return
	}
	if err := config.DB.Create(&notebook).Error; err != nil{
		c.JSON(500, gin.H{"success":false, "message":"Failed to create note"})
		return
	}else{
		c.JSON(201, gin.H{"success":true, "message":"Note added", "note":notebook})
		return
	}

	var existingNotebook models.Notebook
    // Note: If "order" is a reserved keyword in your DB, consider quoting it as needed.
    if err := config.DB.Where("topicID = ? AND \"order\" = ?", notebook.TopicID, notebook.Order).First(&existingNotebook).Error; err == nil {
        // Record found -- duplicate order.
        c.JSON(400, gin.H{"success": false, "message": "Change order number of the title"})
        return
    } else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        c.JSON(500, gin.H{"success": false, "message": "Database error"})
        return
    }

}


func (nc *notebookControllers) GetAllNotes(c *gin.Context){
var notebooks []models.Notebook


}
