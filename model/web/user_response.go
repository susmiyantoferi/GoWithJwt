package web

type UserResponse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}
