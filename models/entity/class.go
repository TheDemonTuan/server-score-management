package entity

import (
	"time"
)

type Class struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"not null;size:100"`
	MaxStudents int    `json:"max_students" gorm:"not null"`

	DepartmentID     uint   `json:"department_id" gorm:"not null;index"`
	HostInstructorID string `json:"host_instructor_id" gorm:"not null;index"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Students []Student `json:"students" gorm:"foreignKey:ClassID"`
}
