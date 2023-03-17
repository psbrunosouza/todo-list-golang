package entities

type User struct {
	Default
	FullName   string      `json:"fullName,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Email      string      `json:"email,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Password   string      `json:"password,omitempty" gorm:"not null;default:null;type:varchar(255)"`
	Profile    string      `json:"profile,omitempty" gorm:"type:varchar(255)"`
	Workspaces []Workspace `json:"workspaces,omitempty"`
}
