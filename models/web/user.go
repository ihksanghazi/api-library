package web

import (
	"time"

	"github.com/google/uuid"
)

type RegisterWebRequest struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Role        string `json:"role" validate:"oneof='member' 'teacher' 'admin' '' "`
	Password    string `json:"password" validate:"required"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	ImageUrl    string `json:"image_url"`
}

type LoginWebRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserWebRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email" validate:"omitempty,email"`
	Password    string `json:"password"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	ImageUrl    string `json:"image_url" validate:"omitempty,url"`
}

type UsersWebResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	ImageUrl    string    `json:"image_url"`
}

type UserWebResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	ImageUrl    string    `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	//association
	Books []BooksWebResponse `gorm:"many2many:borrowings;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:BookID" json:"list_of_books"`
}

func (u *UserWebResponse) TableName() string {
	return "users"
}
