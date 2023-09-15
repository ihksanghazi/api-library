package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	validate := validator.New()
	bookController := controllers.NewBookController(validate)

	r.Post("/", bookController.CreateBookController)

	return r
}
