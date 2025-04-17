package models

import "gorm.io/gorm"

type Notebook struct {
	gorm.Model
	TopicID int `json:"topicID" gorm:"not null"`
	Topic string `json:"topic" gorm:"not null"`
	Title string `json:"title" gorm:"not null;unique"`
	Order int `json:"order" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
	IsPublished bool `json:"isPublished" gorm:"default:false"`
	IsDraft bool `json:"isDraft" gorm:"default:false"`
}
