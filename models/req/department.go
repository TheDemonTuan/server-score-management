package req

type DepartmentCreate struct {
	Name string `json:"name" validate:"required,max=100" `
}

type DepartmentUpdateById struct {
	Name string `json:"name" validate:"required,max=100" `
}
