package models

type Task struct {
	Default
	Description string    `json:"description,omitempty" gorm:"not null;default:null"`
	Name        string    `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	IsDone      bool      `json:"isDone,omitempty" gorm:"default:false;"`
	IsLoop      bool      `json:"isLoop,omitempty" gorm:"default:false;"`
	Workspace   Workspace `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	WorkspaceID int       `json:"workspaceID,omitempty" gorm:"default:null;"`
	Color       Color     `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ColorID     int       `json:"colorID,omitempty" gorm:"default:null;"`
	Rating      Rating    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RatingID    int       `json:"ratingID,omitempty" gorm:"default:null;"`
	Subtasks    []SubTask `json:"subtasks,omitempty"`
}
