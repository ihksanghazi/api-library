package web

import (
	"time"

	"github.com/google/uuid"
)

type CreateBookWebRequest struct {
	Title           string `json:"title" validate:"required"`
	Author          string `json:"author" validate:"required"`
	PublicationYear int    `json:"publicaion_year" validate:"required"`
	ImageUrl        string `json:"image_url"`
	Total           int    `json:"total" validate:"required"`
}

type UpdateBookWebRequest struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	PublicationYear int    `json:"publicaion_year"`
	ImageUrl        string `json:"image_url"`
	Total           int    `json:"total"`
}

type BooksWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publication_year"`
	ImageUrl        string    `json:"image_url"`
	Total           int       `json:"total"`
}

func (b *BooksWebResponse) TableName() string {
	return "books"
}

type BookWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publicaion_year"`
	ImageUrl        string    `json:"image_url"`
	Total           int       `json:"total"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	// Association
	Users []UsersBorrowWebResponse `gorm:"many2many:borrowings;foreignKey:ID;joinForeignKey:BookID;References:ID;joinReferences:UserID" json:"list_of_users"`
}

func (b *BookWebResponse) TableName() string {
	return "books"
}

type BooksBorrowWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear int       `json:"publication_year"`
	ImageUrl        string    `json:"image_url"`
	//association
	Borrow BorrowingWebResponse `gorm:"foreignKey:BookID" json:"borrowing_status"`
}

func (b *BooksBorrowWebResponse) TableName() string {
	return "books"
}
