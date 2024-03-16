package entity

import (
	"time"
)

type Class struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null;size:100"`
	Max  int    `json:"max" gorm:"not null"`

	DepartmentID uint   `json:"department_id" gorm:"not null;index"`
	InstructorID string `json:"instructor_id" gorm:"not null;index"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Students []Student `json:"students" gorm:"foreignKey:ClassID"`
}
