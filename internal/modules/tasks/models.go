package tasks

import (
	"todo-list/internal/modules/workspaces"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Name        string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Workspace   workspaces.Workspace
	WorkspaceID int
}
