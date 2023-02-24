package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null"`
	Title       string `json:"title,omitempty" gorm:"type:varchar(255);not null"`
}
