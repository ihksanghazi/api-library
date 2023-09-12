package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
)

func LoginRouters() *chi.Mux {
	r := chi.NewRouter()

	validator := validator.New()

	authController := controllers.NewAuthController(validator)

	r.Post("/register", authController.Register)

	return r
}
