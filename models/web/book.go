package web

import (
	"github.com/google/uuid"
)

type CreateBookWebRequest struct {
	Title           string `json:"title" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publicaion_year" validate:"required"`
	ImageUrl        string `json:"image_url"`
	Total           int    `json:"total" validate:"required"`
}

type GetAllBooksWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publication_year"`
	ImageUrl        string    `json:"image_url"`
	Total           int       `json:"total"`
}

type BooksWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publication_year"`
	ImageUrl        string    `json:"image_url"`
	//association
	Borrow BorrowingWebResponse `gorm:"foreignKey:BookID" json:"borrowing_status"`
}

func (b *BooksWebResponse) TableName() string {
	return "books"
}
