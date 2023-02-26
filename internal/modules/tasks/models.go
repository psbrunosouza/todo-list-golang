package tasks

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Title       string `json:"title,omitempty" gorm:"not null;default:null;type:varchar(255)"`
}
