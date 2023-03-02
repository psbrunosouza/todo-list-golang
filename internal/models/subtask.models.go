package models

import (
	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Name        string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Task        Task   `json:"task,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TaskID      int    `json:"taskId,omitempty"`
}
