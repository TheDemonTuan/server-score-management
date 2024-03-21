package req

type ClassCreate struct {
	Name             string `json:"name" validate:"required,max=100"`
	MaxStudents      int    `json:"max_students" validate:"required,number,gte=15,lte=60"`
	DepartmentID     uint   `json:"department_id" validate:"required"`
	HostInstructorID string `json:"host_instructor_id" validate:"required"`
}

type ClassUpdateById struct {
	Name             string `json:"name" validate:"required,max=100"`
	MaxStudents      int    `json:"max_students" validate:"required,number,gte=15,lte=60"`
	DepartmentID     uint   `json:"department_id" validate:"required"`
	HostInstructorID string `json:"host_instructor_id" validate:"required"`
}

type ClassDeleteByListId struct {
	ListId []int `json:"list_id" validate:"required,min=1"`
}
