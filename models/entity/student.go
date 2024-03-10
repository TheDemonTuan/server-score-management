package entity

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID              string    `gorm:"primaryKey;size:25"`
	Name            string    `json:"name" gorm:"size:50"`
	BirthDay        time.Time `json:"birth_day"`
	Gender          bool      `json:"gender"`
	Email           string    `json:"email" gorm:"size:50"`
	Phone           string    `json:"phone" gorm:"size:20"`
	Address         string    `json:"address" gorm:"size:100"`
	ClassID         string    `json:"class" gorm:"size:20"`
	Training_System string    `json:"training_system" gorm:"size:30"`
	Cohort          string    `json:"cohort" gorm:"size:20"`

	Transcripts []Transcript   `gorm:"foreignKey:StudentID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
