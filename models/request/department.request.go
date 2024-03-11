package request

type DepartmentCreateRequest struct {
	ID   string `json:"ID" gorm:"primaryKey;size:25"`
	Name string `json:"name" gorm:"size:100"`
}
