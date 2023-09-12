package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/services"
)

func LoginRouters() *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	validator := validator.New()

	authService := services.NewAuthService(ctx)
	authController := controllers.NewAuthController(validator, authService)

	r.Post("/register", authController.RegisterController)

	return r
}
