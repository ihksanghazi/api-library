package routers

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	return r
}
