package models

type SubTask struct {
	Default
	Description string `json:"description,omitempty" gorm:"not null;default:null"`
	Name        string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsDone      bool   `json:"isDone" gorm:"default:false;"`
	Task        Task   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TaskID      int    `json:"taskID,omitempty" gorm:"default:null;"`
}
