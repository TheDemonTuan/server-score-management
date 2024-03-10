package entity

import (
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID           string    `gorm:"primaryKey;size:25"`
	Name         string    `json:"name" gorm:"size:30"`
	DepartmentID string    `json:"department_id" gorm:"size:25"`
	Students     []Student `gorm:"foreignKey:ClassID"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
