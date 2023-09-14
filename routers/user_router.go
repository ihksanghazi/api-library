package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	repo := repositories.Use(db)
	var model domain.User
	userSevice := services.NewUserService(repo, ctx, db, model)
	userController := controllers.NewUserController(userSevice)

	r.Get("/", userController.GetAllUsersController)
	r.Get("/{id}", userController.GetUserByIdController)

	return r
}
