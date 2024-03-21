package req

type CreateClass struct {
	Name             string `json:"name" validate:"required,max=100"`
	MaxStudents      int    `json:"max_students" validate:"required,number,gte=15,lte=60"`
	DepartmentID     uint   `json:"department_id" validate:"required"`
	HostInstructorID string `json:"host_instructor_id" validate:"required"`
}

type UpdateClass struct {
	Name             string `json:"name" validate:"required,max=100"`
	MaxStudents      int    `json:"max_students" validate:"required,number,gte=15,lte=60"`
	DepartmentID     uint   `json:"department_id" validate:"required"`
	HostInstructorID string `json:"host_instructor_id" validate:"required"`
}
