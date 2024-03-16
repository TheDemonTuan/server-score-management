package entity

import (
	"time"
)

type Assignment struct {
	SubjectID    string `json:"subject_id" gorm:"not null;size:50;index"`
	InstructorID string `json:"instructor_id" gorm:"not null;size:25;index"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
