package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/services"
	"github.com/ihksanghazi/api-library/utils"
)

type BookController interface {
	CreateBookController(w http.ResponseWriter, r *http.Request)
	GetAllBookController(w http.ResponseWriter, r *http.Request)
}

type BookControllerImpl struct {
	validate *validator.Validate
	service  services.BookService
}

func NewBookController(validate *validator.Validate, service services.BookService) BookController {
	return &BookControllerImpl{
		validate: validate,
		service:  service,
	}
}

func (b *BookControllerImpl) CreateBookController(w http.ResponseWriter, r *http.Request) {
	// bind req
	var req web.CreateBookWebRequest
	if errBind := json.NewDecoder(r.Body).Decode(&req); errBind != nil {
		utils.ResponseError(w, http.StatusBadRequest, errBind.Error())
		return
	}

	// validate
	if errVal := utils.Validation(b.validate, &req); len(errVal) > 0 {
		utils.ResponseError(w, http.StatusBadRequest, errVal)
		return
	}

	result, errService := b.service.CreateBookService(req)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "OK", result)
}

func (b *BookControllerImpl) GetAllBookController(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, "OK", "Test")
}
