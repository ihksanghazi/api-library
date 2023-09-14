package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/services"
	"github.com/ihksanghazi/api-library/utils"
	"gorm.io/gorm"
)

type UserController interface {
	GetAllUsersController(w http.ResponseWriter, r *http.Request)
	GetUserByIdController(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &UserControllerImpl{
		service: service,
	}
}

func (u *UserControllerImpl) GetAllUsersController(w http.ResponseWriter, r *http.Request) {
	var totalPage, totalLimit int
	// get query parameter
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	if page == "" || limit == "" {
		totalPage = 1
		totalLimit = 5
	} else {
		// parsing to int
		pageInt, errPageInt := strconv.Atoi(page)
		if errPageInt != nil {
			utils.ResponseError(w, http.StatusBadRequest, errPageInt.Error())
			return
		}
		limitInt, errLimitInt := strconv.Atoi(limit)
		if errLimitInt != nil {
			utils.ResponseError(w, http.StatusBadRequest, errPageInt.Error())
			return
		}

		totalPage = pageInt
		totalLimit = limitInt
	}

	users, totalAllPage, errService := u.service.GetAllUserService(totalPage, totalLimit)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	response := web.Pagination{
		Code:        http.StatusOK,
		Status:      "OK",
		CurrentPage: totalPage,
		TotalPage:   totalAllPage,
		Data:        users,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (u *UserControllerImpl) GetUserByIdController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := u.service.GetUserByIdService(id)
	if err == gorm.ErrRecordNotFound {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "OK", user)
}
