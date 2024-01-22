package web

type UserCreateRequest struct {
	Username string `validate:"required,min=1,max=200" json:"username"`
	Email    string `validate:"required,min=1,max=100,email" json:"email"`
	Password string `validate:"required" json:"password"`
}
