package entities

import "time"

type Rating struct {
	Default
	Icon          string    `json:"icon,omitempty" gorm:"not null;default:null"`
	Name          string    `json:"name,omitempty" gorm:"not null;default:null"`
	Value         int       `json:"weight,omitempty" gorm:"not null;default:null"`
	NextIteration time.Time `json:"nextIteration,omitempty" gorm:"not null; default:null;"`
}
