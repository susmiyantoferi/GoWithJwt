package main

import (
	"RestAPIJWT/app"
	controller2 "RestAPIJWT/controller"
	"RestAPIJWT/helper"
	"RestAPIJWT/middleware"
	repository2 "RestAPIJWT/repository"
	service2 "RestAPIJWT/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	//load env
	envErr := godotenv.Load(".env")
	helper.PanicError(envErr)

	db := app.Database()

	validate := validator.New()

	repository := repository2.NewUserRepositoryImpl()

	service := service2.NewUserServiceImpl(repository, db, validate)

	controller := controller2.NewUserControllerImpl(service)

	router := app.RouterNew(controller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
