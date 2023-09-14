package routers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/middleware"
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

	middleware := middleware.NewMiddleware(ctx, db, model)

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
		r.Delete("/logout", authController.LogoutController)
	})

	return r
}
