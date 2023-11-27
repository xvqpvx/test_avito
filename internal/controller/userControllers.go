package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
	"test_avito/internal/helper"
	"test_avito/internal/service"
)

type UserControllers struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserControllers {
	return &UserControllers{UserService: userService}
}

func (uc *UserControllers) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idUser := p.ByName("idUser")
	id, err := strconv.Atoi(idUser)
	helper.PanicIfError(err)

	result := uc.UserService.FindById(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (uc *UserControllers) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := uc.UserService.FindAll(r.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (uc *UserControllers) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userCreateRequest := request.UserCreateRequest{}
	helper.ReadRequestBody(r, &userCreateRequest)

	uc.UserService.Save(r.Context(), userCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	helper.WriteResponseBody(w, webResponse)
}

func (uc *UserControllers) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userUpdateRequest := request.UserUpdateRequest{}
	helper.ReadRequestBody(r, &userUpdateRequest)

	idUser := p.ByName("idUser")
	id, err := strconv.Atoi(idUser)
	helper.PanicIfError(err)
	userUpdateRequest.IdUser = id

	uc.UserService.Update(r.Context(), userUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResponse)
}

func (uc *UserControllers) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idUser := p.ByName("idUser")
	id, err := strconv.Atoi(idUser)
	helper.PanicIfError(err)

	uc.UserService.Delete(r.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	helper.WriteResponseBody(w, webResponse)
}
