package entity

import (
	"time"
)

type Department struct {
	ID   int8   `json:"id" gorm:"not null;primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"not null;size:100"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Classes  []Class   `json:"classes" gorm:"foreignKey:DepartmentID"`
	Students []Student `json:"students" gorm:"foreignKey:DepartmentID"`
	Teachers []Teacher `json:"teachers" gorm:"foreignKey:DepartmentID"`
	Subjects []Subject `json:"subjects" gorm:"foreignKey:DepartmentID"`
}
