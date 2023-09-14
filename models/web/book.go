package web

import (
	"time"

	"github.com/google/uuid"
)

type BooksWebResponse struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	PublicationYear time.Time `json:"publicaion_year"`
	ImageUrl        string    `json:"image_url"`
	Total           int       `json:"total"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	//association
	Borrow BorrowingWebResponse `gorm:"foreignKey:BookID" json:"status_loan"`
}

func (b *BooksWebResponse) TableName() string {
	return "books"
}
