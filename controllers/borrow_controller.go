package controllers

import (
	"net/http"

	"github.com/ihksanghazi/api-library/utils"
)

type BorrowController interface {
	GetBorrowsController(w http.ResponseWriter, r *http.Request)
}

type BorrowControllerImpl struct{}

func NewBorrowController() BorrowController {
	return &BorrowControllerImpl{}
}

func (b *BorrowControllerImpl) GetBorrowsController(w http.ResponseWriter, r *http.Request) {
	utils.ResponseJSON(w, http.StatusOK, "status", "Get Borrow Controller")
}
