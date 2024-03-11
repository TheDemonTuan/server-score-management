package entity

import (
	"gorm.io/gorm"
	"time"
)

type Transcript struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	ProcessScore float64 `json:"process_score"`
	MidtermScore float64 `json:"midterm_score"`
	FinalScore   float64 `json:"final_score"`
	SubjectID    string  `json:"subject_id" gorm:"size:50;index"`
	StudentID    string  `json:"student_id" gorm:"size:25;index"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
