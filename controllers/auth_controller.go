package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/utils"
)

type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
}

type AuthControllerImpl struct {
	validator *validator.Validate
}

func NewAuthController(validator *validator.Validate) AuthController {
	return &AuthControllerImpl{
		validator: validator,
	}
}

func (a *AuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var req web.RegisterWebRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if errVal := utils.Validation(a.validator, req); len(errVal) > 0 {
		utils.ResponseError(w, http.StatusBadRequest, errVal)
		return
	}

	json.NewEncoder(w).Encode(&req)
}
