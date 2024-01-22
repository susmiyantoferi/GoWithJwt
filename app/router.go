package app

import (
	"RestAPIJWT/controller"
	"RestAPIJWT/exception"
	"github.com/julienschmidt/httprouter"
)

func RouterNew(controller controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/users/auth", controller.Auth)
	router.POST("/api/users/refresh-token", controller.RefreshToken)
	router.GET("/api/users", controller.FindAll)
	router.POST("/api/users", controller.Create)
	router.PUT("/api/users/:userId", controller.Update)
	router.DELETE("/api/users/:userId", controller.Delete)
	router.GET("/api/users/:userId", controller.FindById)

	router.PanicHandler = exception.ErrorHandler
	return router

}
