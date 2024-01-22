package controller

import (
	"RestAPIJWT/helper"
	"RestAPIJWT/model/web"
	"RestAPIJWT/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserControllerImpl(service service.UserService) *UserControllerImpl {
	return &UserControllerImpl{Service: service}
}

func (u *UserControllerImpl) Create(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.BodyToRequest(request, &userCreateRequest)

	responseUser := u.Service.Create(request.Context(), userCreateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)
}

func (u *UserControllerImpl) Update(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.BodyToRequest(request, &userUpdateRequest)

	userUpdateRequest.Id = params.ByName("userId")

	responseUser := u.Service.Update(request.Context(), userUpdateRequest)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)

}

func (u *UserControllerImpl) Delete(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	u.Service.Delete(request.Context(), userId)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WritteToBody(writter, response)
}

func (u *UserControllerImpl) FindAll(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	responseUser := u.Service.FindAll(request.Context())
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)
}

func (u *UserControllerImpl) FindById(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	responseUser := u.Service.FindById(request.Context(), userId)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)
}

func (u *UserControllerImpl) Auth(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userAuth := web.UserAuthRequest{}
	helper.BodyToRequest(request, &userAuth)

	responseUser := u.Service.Auth(request.Context(), userAuth)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)
}

func (u *UserControllerImpl) RefreshToken(writter http.ResponseWriter, request *http.Request, params httprouter.Params) {
	token := request.Header.Get("Authorization")

	responseUser := u.Service.RefreshToken(request.Context(), token)
	response := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responseUser,
	}
	helper.WritteToBody(writter, response)
}
