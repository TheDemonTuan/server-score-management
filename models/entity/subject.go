package entity

import (
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID          string `gorm:"primaryKey;size:50"`
	Name        string `gorm:"size:100"`
	Credits     int
	Transcripts []Transcript   `gorm:"foreignKey:SubjectID"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
