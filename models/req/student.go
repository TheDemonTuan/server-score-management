package req

import "time"

type StudentCreate struct {
	FirstName    string    `json:"first_name" validate:"required,min=3,max=50"`
	LastName     string    `json:"last_name" validate:"required,min=3,max=50"`
	Email        string    `json:"email" validate:"required,email,max=100"`
	Address      string    `json:"address" validate:"required,max=100"`
	BirthDay     time.Time `json:"birth_day" validate:"required"`
	Phone        string    `json:"phone" validate:"required,max=20"`
	Gender       bool      `json:"gender" validate:"required,boolean"`
	DepartmentID uint      `json:"department_id" validate:"required"`
	ClassID      uint      `json:"class_id" validate:"required"`
}

type StudentUpdateById struct {
	FirstName    string    `json:"first_name" validate:"required,min=3,max=50"`
	LastName     string    `json:"last_name" validate:"required,min=3,max=50"`
	Email        string    `json:"email" validate:"required,email,max=100"`
	Address      string    `json:"address" validate:"required,max=100"`
	BirthDay     time.Time `json:"birth_day" validate:"required"`
	Phone        string    `json:"phone" validate:"required,max=20"`
	Gender       bool      `json:"gender" validate:"required,boolean"`
	DepartmentID uint      `json:"department_id" validate:"required"`
	ClassID      uint      `json:"class_id" validate:"required"`
}

type StudentDeleteByListId struct {
	ListId []int `json:"list_id" validate:"required,min=1"`
}
