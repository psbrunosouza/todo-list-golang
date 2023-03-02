package models

import (
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	Name  string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Tasks []Task
}
