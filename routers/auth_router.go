package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	var model domain.User
	validator := validator.New()
	authService := services.NewAuthService(ctx, db, model)
	authController := controllers.NewAuthController(validator, authService)

	r.Delete("/logout", authController.LogoutController)
	r.Get("/token", authController.GetTokenController)
	r.Post("/register", authController.RegisterController)
	r.Post("/login", authController.LoginController)

	return r
}
