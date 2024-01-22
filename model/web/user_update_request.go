package web

type UserUpdateRequest struct {
	Id       string `validate:"required" json:"id"`
	Username string `validate:"required,min=1,max=200" json:"username"`
}
