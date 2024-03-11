package req

type DepartmentCreate struct {
	ID   int8   `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,max=100" `
}

type DepartmentUpdate struct {
	Name string `json:"name" validate:"required,max=100" `
}
