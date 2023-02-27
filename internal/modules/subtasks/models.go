package subtasks

import (
	"todo-list/internal/modules/tasks"

	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Title       string `json:"title,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Task        tasks.Task
}
