package routers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/middleware"
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

	middleware := middleware.NewMiddleware(repo)

	r.Group(func(r chi.Router) {
		r.Use(middleware.ValidToken)
		r.Use(middleware.IsAdmin)
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Berhasil Access"))
		})
	})

	r.Group(func(r chi.Router) {
		r.Post("/register", authController.RegisterController)
		r.Post("/login", authController.LoginController)
		r.Get("/token", authController.GetTokenController)
	})

	return r
}
