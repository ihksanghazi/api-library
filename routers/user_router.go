package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	repo := repositories.Use(db)
	userSevice := services.NewUserService(repo)
	userController := controllers.NewUserController(userSevice)

	r.Get("/", userController.GetAllUsersController)

	return r
}
