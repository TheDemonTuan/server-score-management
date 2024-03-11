package entity

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID           string    `gorm:"primaryKey;size:25"`
	FirstName    string    `json:"first_name" gorm:"size:50"`
	LastName     string    `json:"last_name" gorm:"size:50"`
	BirthDay     time.Time `json:"birth_day"`
	Gender       bool      `json:"gender"`
	Email        string    `json:"email" gorm:"unique;size:50"`
	Phone        string    `json:"phone" gorm:"unique;size:20"`
	Address      string    `json:"address" gorm:"size:100"`
	ClassID      int8      `json:"class_id" gorm:"not null;index"`
	DepartmentID int8      `json:"department_id" gorm:"not null;index"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Transcripts []Transcript `gorm:"foreignKey:StudentID"`
}
