package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ihksanghazi/api-library/database"
	"github.com/ihksanghazi/api-library/routers"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect database
	db := database.ConnectDB()

	// migration
	// db.AutoMigrate(domain.User{}, domain.Book{}, domain.Borrowing{})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/auth", routers.AuthRouter(db))
	r.Mount("/api/user", routers.UserRouter(db))
	r.Mount("/api/book", routers.BookRouter(db))
	r.Mount("/api/borrow", routers.BorrowRouter(db))

	http.ListenAndServe(":3000", r)
}
