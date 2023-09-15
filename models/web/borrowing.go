package web

import (
	"time"

	"github.com/google/uuid"
)

type BorrowingWebResponse struct {
	BookID        uuid.UUID `json:"-"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturnDate    time.Time `json:"return_date"`
	Status        string    `json:"status"`
}

func (b *BorrowingWebResponse) TableName() string {
	return "borrowings"
}
