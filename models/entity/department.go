package entity

import (
	"gorm.io/gorm"
	"time"
)

type Department struct {
	ID   int8   `json:"id" gorm:"not null;primaryKey"`
	Name string `json:"name" gorm:"not null;size:100"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Classes  []Class   `gorm:"foreignKey:DepartmentID"`
	Students []Student `gorm:"foreignKey:DepartmentID"`
	Teachers []Teacher `gorm:"foreignKey:DepartmentID"`
	Subjects []Subject `gorm:"foreignKey:DepartmentID"`
}
