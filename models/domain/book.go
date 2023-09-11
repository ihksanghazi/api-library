package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID              uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title           string         `gorm:"not null" json:"title"`
	Author          string         `gorm:"not null" json:"author"`
	PublicationYear time.Time      `json:"publicaion_year"`
	ImageUrl        string         `json:"image_url"`
	Total           int            `json:"total"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
