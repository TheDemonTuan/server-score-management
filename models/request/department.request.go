package request

type DepartmentCreateRequest struct {
	ID   string `json:"id" validate:"required,max=25"`
	Name string `json:"name" validate:"required,max=100" `
}

type DepartmentUpdateRequest struct {
	Name string `json:"name" validate:"required,max=100" `
}
