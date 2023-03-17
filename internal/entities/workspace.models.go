package entities

type Workspace struct {
	Default
	Name    string `json:"name,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Tasks   []Task `json:"tasks,omitempty"`
	User    User   `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID  int    `json:"userID,omitempty" gorm:"default:null;"`
	Color   Color  `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ColorID int    `json:"colorID,omitempty" gorm:"default:null;"`
}
