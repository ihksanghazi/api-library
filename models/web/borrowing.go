package web

import (
	"time"

	"github.com/google/uuid"
)

type BorrowsWebResponse struct {
	User          UsersWebResponse `gorm:"foreignKey:UserID" json:"user"`
	Book          BooksWebResponse `gorm:"foreignKey:BookID" json:"book"`
	BookID        uuid.UUID        `json:"-"`
	UserID        uuid.UUID        `json:"-"`
	BorrowingDate time.Time        `json:"borrowing_date"`
	ReturnDate    time.Time        `json:"return_date"`
	Status        string           `json:"status"`
}

func (b *BorrowsWebResponse) TableName() string {
	return "borrowings"
}

type BorrowingWebResponse struct {
	BookID        uuid.UUID `json:"-"`
	UserID        uuid.UUID `json:"-"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturnDate    time.Time `json:"return_date"`
	Status        string    `json:"status"`
}

func (b *BorrowingWebResponse) TableName() string {
	return "borrowings"
}
