package utils

import (
	"RestAPIJWT/exception"
	"RestAPIJWT/helper"
	"RestAPIJWT/model/web"
	jwt "github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(request web.TokenCreaterequest, value time.Duration) string {
	var jwtSecretToken = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	jwtExpiredTime := time.Now().Add(time.Minute * value)

	claimToken := &web.TokenClaims{
		UserId:   request.UserId,
		Username: request.Username,
		Email:    request.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(jwtExpiredTime),
		},
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, claimToken)
	tokenStr, err := token.SignedString(jwtSecretToken)
	helper.PanicError(err)

	return tokenStr
}

func TokenClaims(userToken string) web.TokenClaims {
	var jwtSecretToken = []byte(os.Getenv("JWT_TOKEN_SECRET"))
	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(userToken, claims,
		func(j *jwt.Token) (interface{}, error) {
			return jwtSecretToken, nil
		},
	)
	if err != nil {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}
	if !token.Valid {
		panic(exception.NewErrorUnauthorized(err.Error()))
	}
	return *claims
}
