package models

import "time"

type Iteration struct {
	Default
	Rating        Rating    `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RatingID      int       `json:"ratingID,omitempty" gorm:"default:null;"`
	Task          Task      `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TaskID        int       `json:"taskID,omitempty" gorm:"default:null;"`
	NextIteration time.Time `json:"nextIteration,omitempty" gorm:"not null; default:null;"`
}
