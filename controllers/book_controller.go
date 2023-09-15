package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/services"
	"github.com/ihksanghazi/api-library/utils"
)

type BookController interface {
	CreateBookController(w http.ResponseWriter, r *http.Request)
	GetAllBookController(w http.ResponseWriter, r *http.Request)
	UpdateBookController(w http.ResponseWriter, r *http.Request)
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
	var totalPage, totalLimit int
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

	result, totalAllPage, errService := b.service.GetAllBookService(totalPage, totalLimit)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	response := web.Pagination{
		Code:        http.StatusOK,
		Status:      "OK",
		CurrentPage: totalPage,
		TotalPage:   totalAllPage,
		Data:        result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (b *BookControllerImpl) UpdateBookController(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// bind req
	var req web.UpdateBookWebRequest
	if errBind := json.NewDecoder(r.Body).Decode(&req); errBind != nil {
		utils.ResponseError(w, http.StatusBadRequest, errBind.Error())
		return
	}

	result, errService := b.service.UpdateBookService(id, req)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "OK", result)
}
