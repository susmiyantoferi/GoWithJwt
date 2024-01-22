package middleware

import (
	"RestAPIJWT/helper"
	"RestAPIJWT/model/web"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{handler: handler}

}

func (middleware *AuthMiddleware) unautorized(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusUnauthorized)

	response := web.WebResponse{
		Code:   http.StatusUnauthorized,
		Status: "UNAUTHORIZED",
	}
	helper.WritteToBody(writer, response)
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" && (request.RequestURI == "/api/users" || request.RequestURI == "/api/users/auth") {
		middleware.handler.ServeHTTP(writer, request)
	} else {
		tokenAuth := request.Header.Get("Authorization")
		if tokenAuth == "" {
			middleware.unautorized(writer, request)
			return
		}

		var jwtSecretToken = []byte(os.Getenv("JWT_TOKEN_SECRET"))
		claims := &web.TokenClaims{}

		token, err := jwt.ParseWithClaims(tokenAuth, claims,
			func(j *jwt.Token) (interface{}, error) {
				return jwtSecretToken, nil
			},
		)
		if err != nil {
			middleware.unautorized(writer, request)
			return
		}
		if !token.Valid {
			middleware.unautorized(writer, request)
			return
		}

		middleware.handler.ServeHTTP(writer, request)

	}
}
