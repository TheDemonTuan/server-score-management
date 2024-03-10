package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserName  string         `json:"user_name" gorm:"primaryKey;size:50"`
	Password  string         `json:"password" gorm:"size:50"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
