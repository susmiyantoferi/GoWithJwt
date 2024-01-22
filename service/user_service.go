package service

import (
	"RestAPIJWT/model/web"
	"context"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, userId string)
	FindAll(ctx context.Context) []web.UserResponse
	FindById(ctx context.Context, userId string) web.UserResponse
	Auth(ctx context.Context, request web.UserAuthRequest) web.TokenResponse
	RefreshToken(ctx context.Context, refreshToken string) web.TokenResponse
}
