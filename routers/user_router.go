package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	repo := repositories.Use(db)
	userSevice := services.NewUserService(repo, ctx)
	userController := controllers.NewUserController(userSevice)

	r.Get("/", userController.GetAllUsersController)
	r.Get("/{id}", userController.GetUserByIdController)

	return r
}
