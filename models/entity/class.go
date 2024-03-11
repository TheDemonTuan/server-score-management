package entity

type Class struct {
	ID           int8   `json:"id" gorm:"not null;primaryKey"`
	Name         string `json:"name" gorm:"not null;size:100"`
	DepartmentID int8   `json:"department_id" gorm:"not null;index"`

	Students []Student `gorm:"foreignKey:ClassID"`
}
