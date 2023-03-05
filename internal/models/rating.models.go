package models

type Rating struct {
	Default
	Icon string `json:"icon,omitempty" gorm:"not null;default:null"`
	Name string `json:"name,omitempty" gorm:"not null;default:null"`
}
