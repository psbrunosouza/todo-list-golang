package models

type Workspace struct {
	Default
	Name   string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Tasks  []Task `json:"tasks,omitempty"`
	User   User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID int    `json:"userID,omitempty" gorm:"default:null;"`
}
