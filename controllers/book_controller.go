package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/utils"
)

type BookController interface {
	CreateBookController(w http.ResponseWriter, r *http.Request)
}

type BookControllerImpl struct {
	validate *validator.Validate
}

func NewBookController(validate *validator.Validate) BookController {
	return &BookControllerImpl{
		validate: validate,
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

	utils.ResponseJSON(w, http.StatusOK, "OK", req)
}
