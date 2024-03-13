package entity

import (
	"time"
)

type Department struct {
	ID   int8   `json:"id" gorm:"not null;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null;size:100"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Classes  []Class   `gorm:"foreignKey:DepartmentID"`
	Students []Student `gorm:"foreignKey:DepartmentID"`
	Teachers []Teacher `gorm:"foreignKey:DepartmentID"`
	Subjects []Subject `gorm:"foreignKey:DepartmentID"`
}
