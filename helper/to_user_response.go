package helper

import (
	"RestAPIJWT/model"
	"RestAPIJWT/model/web"
)

func ToUserResponse(user model.User) web.UserResponse {
	return web.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		CreateAt: user.CreateAt,
		UpdateAt: user.UpdateAt,
	}
}
