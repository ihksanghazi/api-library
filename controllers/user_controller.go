package controllers

import (
	"net/http"

	"github.com/ihksanghazi/api-library/utils"
)

type UserController interface {
	GetAllUsersController(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
}

func NewUserController() UserController {
	return &UserControllerImpl{}
}

func (u *UserControllerImpl) GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, "OK", "Test")
}
