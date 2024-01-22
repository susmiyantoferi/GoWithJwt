package web

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
