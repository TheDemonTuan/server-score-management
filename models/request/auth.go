package request

type AuthLogin struct {
	UserName string `json:"user_name" `
	Password string `json:"password"`
}

type AuthRegister struct {
	FirstName string `json:"first_name" validate:"required,min=3,max=50"`
	LastName  string `json:"last_name" validate:"required,min=3,max=50"`
	UserName  string `json:"username" validate:"required,min=3,max=30"`
	Password  string `json:"password" validate:"required,min=8"`
}
