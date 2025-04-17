package models

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name string `json:"name"`
	Company string `json:"company"`
}