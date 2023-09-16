package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/services"
	"github.com/ihksanghazi/api-library/utils"
)

type BorrowController interface {
	GetBorrowsController(w http.ResponseWriter, r *http.Request)
}

type BorrowControllerImpl struct {
	service services.BorrowService
}

func NewBorrowController(service services.BorrowService) BorrowController {
	return &BorrowControllerImpl{
		service: service,
	}
}

func (b *BorrowControllerImpl) GetBorrowsController(w http.ResponseWriter, r *http.Request) {
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

	borrows, totalAllPage, errService := b.service.GetBorrowsService(totalPage, totalLimit)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	response := web.Pagination{
		Code:        http.StatusOK,
		Status:      "OK",
		CurrentPage: totalPage,
		TotalPage:   totalAllPage,
		Data:        borrows,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}
