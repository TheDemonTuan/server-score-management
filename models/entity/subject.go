package entity

import (
	"time"
)

type Subject struct {
	ID                string `json:"id" gorm:"primaryKey;size:50"`
	Name              string `json:"name" gorm:"not null;size:100"`
	Credits           int8   `json:"credits" gorm:"not null"`
	ProcessPercentage int8   `json:"process_percentage" gorm:"not null"`
	MidtermPercentage int8   `json:"midterm_percentage" gorm:"not null"`
	FinalPercentage   int8   `json:"final_percentage" gorm:"not null"`

	DepartmentID uint `json:"department_id" gorm:"not null;index"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Grades      []Grade      `json:"grades" gorm:"foreignKey:SubjectID"`
	Assignments []Assignment `json:"assignments" gorm:"foreignKey:SubjectID"`
}
