package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/middleware"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	var model domain.User
	validation := validator.New()

	userSevice := services.NewUserService(ctx, db, model)
	userController := controllers.NewUserController(userSevice, validation)

	authMiddleware := middleware.NewMiddleware(ctx, db, model)

	// admin
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.ValidToken)
		r.Use(authMiddleware.IsAdmin)
		r.Get("/", userController.GetAllUsersController)
		r.Get("/{id}", userController.GetUserByIdController)
	})

	// user
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.ValidToken)
		r.Put("/{id}", userController.UpdateUserController)
		r.Delete("/{id}", userController.DeleteUserController)
	})

	return r
}
