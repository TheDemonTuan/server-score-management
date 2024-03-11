package entity

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID      string  `json:"ID" gorm:"primaryKey;size:25"`
	Name    string  `json:"name" gorm:"size:100"`
	Classes []Class `json:"Classes" gorm:"foreignKey:DepartmentID"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
