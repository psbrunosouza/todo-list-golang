package entities

type IterationSetting struct {
	Default
	TotalIterations int `json:"totalIterations,omitempty" gorm:"default:3"`
}
