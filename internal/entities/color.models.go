package entities

type Color struct {
	Default
	Hex  string `json:"hex,omitempty" gorm:"not null;default:null"`
	Name string `json:"name,omitempty" gorm:"not null;default:null"`
}
