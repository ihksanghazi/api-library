package web

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateBookWebRequest struct {
	Title           string         `gorm:"not null" json:"title"`
	Author          string         `gorm:"not null" json:"author"`
	PublicationYear int            `json:"publicaion_year"`
	ImageUrl        string         `json:"image_url"`
	Total           int            `json:"total"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
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
