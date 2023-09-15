package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func BookRouter(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Book"))
	})

	return r
}
