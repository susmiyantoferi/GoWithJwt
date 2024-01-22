package service

import (
	"RestAPIJWT/exception"
	"RestAPIJWT/helper"
	"RestAPIJWT/model"
	"RestAPIJWT/model/web"
	"RestAPIJWT/repository"
	"RestAPIJWT/utils"
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	err := u.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	hashPassword, err := utils.HashPassword(request.Password)
	helper.PanicError(err)

	user := model.User{
		Id:       utils.Uuid(),
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}
	user = u.UserRepository.Create(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (u *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := u.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	//find id
	user, err := u.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	//update data
	user.Username = request.Username
	user.UpdateAt = time.Now().Unix()

	user = u.UserRepository.Update(ctx, tx, user)
	return helper.ToUserResponse(user)
}

func (u *UserServiceImpl) Delete(ctx context.Context, userId string) {
	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	//find id
	user, err := u.UserRepository.FindById(ctx, tx, userId)
	helper.PanicError(err)
	//delete data
	u.UserRepository.Delete(ctx, tx, user)
}

func (u *UserServiceImpl) FindAll(ctx context.Context) []web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	users := u.UserRepository.FindAll(ctx, tx)

	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, helper.ToUserResponse(user))
	}
	return userResponses
}

func (u *UserServiceImpl) FindById(ctx context.Context, userId string) web.UserResponse {
	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	//find id
	user, err := u.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToUserResponse(user)
}

func (u *UserServiceImpl) Auth(ctx context.Context, request web.UserAuthRequest) web.TokenResponse {
	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	user, err := u.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}

	jwtTokenExpired, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRED"))
	jwtTokenRefresh, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_REFRESH"))

	tokenCreateRequest := web.TokenCreaterequest{
		UserId:   user.Id,
		Username: user.Username,
		Email:    user.Email,
	}

	token := web.TokenResponse{
		Token: utils.CreateToken(
			tokenCreateRequest,
			time.Duration(jwtTokenExpired),
		),
		TokenRefresh: utils.CreateToken(
			tokenCreateRequest,
			time.Duration(jwtTokenRefresh),
		),
	}

	return token
}

func (u *UserServiceImpl) RefreshToken(ctx context.Context, refreshToken string) web.TokenResponse {
	tx, err := u.DB.Begin()
	helper.PanicError(err)
	defer helper.DeferDb(tx)

	tokenClaims := utils.TokenClaims(refreshToken)
	_, err = u.UserRepository.FindById(ctx, tx, tokenClaims.UserId)
	if err != nil {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}

	tokenCreateRequest := web.TokenCreaterequest{
		UserId:   tokenClaims.UserId,
		Username: tokenClaims.Username,
		Email:    tokenClaims.Email,
	}

	jwtTokenExpired, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_EXPIRED"))
	jwtTokenRefresh, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_REFRESH"))

	token := web.TokenResponse{
		Token: utils.CreateToken(
			tokenCreateRequest,
			time.Duration(jwtTokenExpired),
		),
		TokenRefresh: utils.CreateToken(
			tokenCreateRequest,
			time.Duration(jwtTokenRefresh),
		),
	}

	return token

}
