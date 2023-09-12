package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ihksanghazi/api-library/database"
	"github.com/ihksanghazi/api-library/repositories"
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
	repositories.SetDefault(db)

	// migration
	// database.DB.AutoMigrate(domain.User{}, domain.Book{}, domain.Borrowing{})

	//generating query in folder repositories
	// database.GenerateQuery(database.DB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/api/auth", routers.LoginRouters(db))

	http.ListenAndServe(":3000", r)
}
