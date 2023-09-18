package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ihksanghazi/api-library/database"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/routers"
	"github.com/ihksanghazi/api-library/services"
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

	http.ListenAndServe(":3000", r)

	// Create ticker for check every day
	ticker := time.NewTicker(24 * time.Hour)

	var ctx context.Context
	var book domain.Book
	var borrow domain.Borrowing
	bookService := services.NewBookService(db, ctx, book, borrow)

	for range ticker.C {
		// Update Expired Service Every Day
		bookService.UpdateExpiredService()
	}

}
