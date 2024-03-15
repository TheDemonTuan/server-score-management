package entity

import (
	"gorm.io/gorm"
	"time"
)

type Subject struct {
	ID                string `json:"id" gorm:"not null;primaryKey;size:50"`
	Name              string `json:"name" gorm:"not null;size:100"`
	Credits           int    `json:"credits" gorm:"not null"`
	ProcessPercentage int8   `json:"process_percentage"`
	MidtermPercentage int8   `json:"midterm_percentage" gorm:"not null"`
	FinalPercentage   int8   `json:"final_percentage" gorm:"not null"`
	DepartmentID      int8   `json:"department_id" gorm:"not null;index"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Transcripts []Transcript `json:"transcripts" gorm:"foreignKey:SubjectID"`
}
