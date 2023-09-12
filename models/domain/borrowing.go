package domain

import (
	"time"

	"github.com/google/uuid"
)

type Borrowing struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	BookID     uuid.UUID `gorm:"type:uuid;foreignKey" json:"book_id"`
	UserID     uuid.UUID `gorm:"type:uuid;foreignKey" json:"user_id"`
	LoanDate   time.Time `json:"loan_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
}
