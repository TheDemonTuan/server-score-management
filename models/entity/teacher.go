package entity

import (
	"gorm.io/gorm"
	"time"
)

type Teacher struct {
	ID       string    `gorm:"primaryKey;size:25"`
	Name     string    `json:"name" gorm:"size:50"`
	BirthDay time.Time `json:"birth_day"`
	Email    string    `json:"email" gorm:"size:50"`
	Phone    string    `json:"phone" gorm:"size:20"`
	Address  string    `json:"address" gorm:"size:100"`
	Degree   string    `json:"degree" gorm:"size:50"`

	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Subjects []Subject `gorm:"many2many:teacher_subjects"`
}
