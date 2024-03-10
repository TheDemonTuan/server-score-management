package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transcript struct {
	ID                uuid.UUID      `gorm:"type:char(36);primaryKey"`
	ProcessScore      float64        `json:"process_score"`
	MidtermScore      float64        `json:"midterm_score"`
	FinalScore        float64        `json:"final_score"`
	FinalGrade        float64        `json:"final_grade"`
	ProcessPercentage float64        `json:"process_percentage"`
	MidtermPercentage float64        `json:"midterm_percentage"`
	FinalPercentage   float64        `json:"final_percentage"`
	SubjectID         string         `json:"subject_id" gorm:"size:50"`
	StudentID         string         `json:"student_id" gorm:"size:25"`
	CreatedAt         time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
