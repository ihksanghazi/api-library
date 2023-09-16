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

func BookRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	validate := validator.New()
	var book domain.Book
	var borrow domain.Borrowing

	bookService := services.NewBookService(db, ctx, book, borrow)
	bookController := controllers.NewBookController(validate, bookService)

	r.Get("/", bookController.GetAllBookController)
	r.Get("/{id}", bookController.GetBookByIdController)
	r.Get("/borrow/{id}", bookController.BorrowBookController)
	r.Post("/", bookController.CreateBookController)
	r.Put("/{id}", bookController.UpdateBookController)
	r.Delete("/{id}", bookController.DeleteBookController)

	return r
}
