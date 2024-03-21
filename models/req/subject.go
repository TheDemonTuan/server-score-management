package req

type SubjectCreate struct {
	Name              string `json:"name" validate:"required,max=100"`
	Credits           int8   `json:"credits" validate:"required"`
	ProcessPercentage int8   `json:"process_percentage" validate:"gte=0,lte=50"`
	MidtermPercentage int8   `json:"midterm_percentage" validate:"gte=0,lte=50"`
	FinalPercentage   int8   `json:"final_percentage" validate:"required,gte=50,lte=100"`
	DepartmentID      uint   `json:"department_id" validate:"required"`
}

type SubjectUpdateById struct {
	Name              string `json:"name" validate:"required,max=100"`
	Credits           int8   `json:"credits" validate:"required"`
	ProcessPercentage int8   `json:"process_percentage" validate:"required,gte=0,lte=50"`
	MidtermPercentage int8   `json:"midterm_percentage" validate:"required,gte=0,lte=50"`
	FinalPercentage   int8   `json:"final_percentage" validate:"required,gte=50,lte=100"`
	DepartmentID      uint   `json:"department_id" validate:"required"`
}
