package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	validate := validator.New()
	bookService := services.NewBookService(db, ctx)
	bookController := controllers.NewBookController(validate, bookService)

	r.Post("/", bookController.CreateBookController)

	return r
}
