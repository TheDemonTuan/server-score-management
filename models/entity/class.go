package entity

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID           int8   `json:"id" gorm:"not null;primaryKey"`
	Name         string `json:"name" gorm:"not null;size:100"`
	DepartmentID int8   `json:"department_id" gorm:"not null;index"`
	TeacherID    string `json:"teacher_id" gorm:"not null;index"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Students []Student `gorm:"foreignKey:ClassID"`
}
