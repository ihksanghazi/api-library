package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/services"
	"github.com/ihksanghazi/api-library/utils"
)

type AuthController interface {
	RegisterController(w http.ResponseWriter, r *http.Request)
	LoginController(w http.ResponseWriter, r *http.Request)
	GetTokenController(w http.ResponseWriter, r *http.Request)
	LogoutController(w http.ResponseWriter, r *http.Request)
}

type AuthControllerImpl struct {
	validator *validator.Validate
	service   services.AuthService
}

func NewAuthController(validator *validator.Validate, service services.AuthService) AuthController {
	return &AuthControllerImpl{
		validator: validator,
		service:   service,
	}
}

func (a *AuthControllerImpl) RegisterController(w http.ResponseWriter, r *http.Request) {
	var req web.RegisterWebRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// cek validation
	if errVal := utils.Validation(a.validator, req); len(errVal) > 0 {
		utils.ResponseError(w, http.StatusBadRequest, errVal)
		return
	}

	user, errSer := a.service.RegisterService(&req)
	if errSer != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errSer.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "OK", user)
}

func (a *AuthControllerImpl) LoginController(w http.ResponseWriter, r *http.Request) {
	//bind json
	var req web.LoginWebRequest
	if errBind := json.NewDecoder(r.Body).Decode(&req); errBind != nil {
		utils.ResponseError(w, http.StatusBadRequest, errBind.Error())
		return
	}

	// validation
	if errValidation := utils.Validation(a.validator, req); len(errValidation) > 0 {
		utils.ResponseError(w, http.StatusBadRequest, errValidation)
		return
	}

	timeRefreshToken := time.Now().Add(time.Hour * 24)
	timeAccessToken := time.Now().Add(time.Second * 30)

	accessToken, refreshToken, errService := a.service.LoginService(&req, &timeRefreshToken, &timeAccessToken)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	// set refresh token to cookies
	cookie := http.Cookie{
		Name:     "AccessToken",
		HttpOnly: true,
		Value:    *refreshToken,
		SameSite: http.SameSiteNoneMode,
		Expires:  timeRefreshToken,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	// utils
	utils.ResponseJSON(w, http.StatusOK, "Your Access Token", accessToken)
}

func (a *AuthControllerImpl) GetTokenController(w http.ResponseWriter, r *http.Request) {
	refreshToken, errCookie := r.Cookie("AccessToken")
	if errCookie != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, errService := a.service.GetTokenService(&refreshToken.Value)
	if errService != nil {
		utils.ResponseError(w, http.StatusInternalServerError, errService.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusOK, "Your Access Token", accessToken)
}

func (a *AuthControllerImpl) LogoutController(w http.ResponseWriter, r *http.Request) {
	// set cookie to -1
	cookie := http.Cookie{
		Name:     "AccessToken",
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		MaxAge:   -1,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	utils.ResponseJSON(w, http.StatusOK, "You Have Logged Out", nil)
}
