package req

type DepartmentCreate struct {
	Name string `json:"name" validate:"required,min=3,max=100" `
}

type DepartmentUpdateById struct {
	Name string `json:"name" validate:"required,min=3,max=100" `
}
