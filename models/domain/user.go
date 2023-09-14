package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Username     string         `gorm:"not null" json:"username"`
	Email        string         `gorm:"unique;not null" json:"email"`
	Role         string         `gorm:"not null" json:"role"`
	Password     string         `gorm:"not null" json:"password"`
	RefreshToken string         `json:"refresh_token"`
	Address      string         `json:"address"`
	PhoneNumber  string         `json:"phone_number"`
	ImageUrl     string         `json:"image_url"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	//association
	Borrowing Borrowing `gorm:"foreignKey:UserID"`
	Books     []Book    `gorm:"many2many:borrowings;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:BookID"`
}
