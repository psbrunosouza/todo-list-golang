package models

type Task struct {
	Default
	Description string    `json:"description,omitempty" gorm:"not null;default:null"`
	Name        string    `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Workspace   Workspace `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WorkspaceID int       `json:"workspaceID,omitempty"`
	Subtasks    []SubTask `json:"subtasks,omitempty"`
}
