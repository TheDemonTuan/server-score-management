package req

import "time"

type TeacherCreate struct {
	ID           string    `json:"id" validate:"required" `
	Name         string    `json:"name" validate:"required,max=50"`
	BirthDay     time.Time `json:"birth_day" validate:"required" `
	Email        string    `json:"email" validate:"required,max=50, email"`
	Phone        string    `json:"phone" validate:"required,max=20, phone"`
	Address      string    `json:"address" validate:"required,max=100"`
	Degree       string    `json:"degree" validate:"required,max=50"`
	DepartmentID int8      `json:"department_id" validate:"required"`
}

type TeacherUpdate struct {
	Name         string    `json:"name" validate:"required,max=50"`
	BirthDay     time.Time `json:"birth_day" validate:"required" `
	Email        string    `json:"email" validate:"required,max=50, email"`
	Phone        string    `json:"phone" validate:"required,max=20, phone"`
	Address      string    `json:"address" validate:"required,max=100"`
	Degree       string    `json:"degree" validate:"required,max=50"`
	DepartmentID int8      `json:"department_id" validate:"required"`
}
