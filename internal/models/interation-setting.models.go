package models

type IterationSetting struct {
	Default
	TotalIterations int `json:"totalIterations,omitempty" gorm:"default:3"`
}
