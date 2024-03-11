package req

type DepartmentCreate struct {
	ID   string `json:"id" validate:"required,max=25"`
	Name string `json:"name" validate:"required,max=100" `
}

type DepartmentUpdate struct {
	Name string `json:"name" validate:"required,max=100" `
}
