package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ihksanghazi/api-library/controllers"
	"gorm.io/gorm"
)

func BorrowRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	borrowController := controllers.NewBorrowController()

	r.Get("/", borrowController.GetBorrowsController)

	return r
}
