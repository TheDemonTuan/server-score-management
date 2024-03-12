package req

type SubjectCreate struct {
	//ID                string `json:"id" validate:"required" `
	Name              string `json:"name" validate:"required,max=100"`
	Credits           int    `json:"credits" validate:"required"`
	ProcessPercentage int8   `json:"process_percentage" validate:""`
	MidtermPercentage int8   `json:"midterm_percentage" validate:"required"`
	FinalPercentage   int8   `json:"final_percentage" validate:"required"`
	DepartmentID      int8   `json:"department_id" validate:"required"`
}

type SubjectUpdate struct {
	Name              string `json:"name" validate:"required,max=100"`
	Credits           int    `json:"credits" validate:"required, max = 14 , min = 2"`
	ProcessPercentage int8   `json:"process_percentage" validate:"required, max=100, min=0"`
	MidtermPercentage int8   `json:"midterm_percentage" validate:"required, max=100, min=0"`
	FinalPercentage   int8   `json:"final_percentage" validate:"required , max=100, min=0"`
	DepartmentID      int8   `json:"department_id" validate:"required"`
}
