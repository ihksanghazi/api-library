package routers

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/services"
	"gorm.io/gorm"
)

func BorrowRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	var ctx context.Context
	var model domain.Borrowing

	borrowService := services.NewBorrowService(db, ctx, model)
	borrowController := controllers.NewBorrowController(borrowService)

	r.Get("/", borrowController.GetBorrowsController)

	return r
}
