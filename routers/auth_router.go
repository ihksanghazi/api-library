package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func LoginRouters(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	validator := validator.New()
	repo := repositories.Use(db)
	authService := services.NewAuthService(ctx, repo)
	authController := controllers.NewAuthController(validator, authService)

	r.Post("/register", authController.RegisterController)
	r.Post("/login", authController.LoginController)

	return r
}
