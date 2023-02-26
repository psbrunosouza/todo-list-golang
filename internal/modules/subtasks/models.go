package subtasks

import (
	task_models "todo-list/internal/modules/tasks/models"

	"gorm.io/gorm"
)

type SubTask struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Title       string `json:"title,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Task        task_models.Task
}
