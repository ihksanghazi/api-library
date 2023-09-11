package routers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func LoginRouters() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}
