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

func BookRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	validate := validator.New()
	var book domain.Book
	var borrow domain.Borrowing
	var user domain.User

	bookService := services.NewBookService(db, ctx, book, borrow)
	bookController := controllers.NewBookController(validate, bookService)

	authMiddleware := middleware.NewMiddleware(ctx, db, user)

	//admin
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.ValidToken)
		r.Use(authMiddleware.IsAdmin)
		r.Post("/", bookController.CreateBookController)
		r.Put("/{id}", bookController.UpdateBookController)
		r.Delete("/{id}", bookController.DeleteBookController)
		r.Get("/expired", bookController.GetAllExpiredController)
	})

	//user
	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.ValidToken)
		r.Get("/", bookController.GetAllBookController)
		r.Get("/{id}", bookController.GetBookByIdController)
		r.Get("/borrow/{id}", bookController.BorrowBookController)
		r.Get("/return/{id}", bookController.ReturnBookController)
	})

	return r
}
