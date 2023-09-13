package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	userController := controllers.NewUserController()

	r.Get("/", userController.GetAllUsersController)

	return r
}
