package web

import "github.com/google/uuid"

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

type UsersWebResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	ImageUrl    string    `json:"image_url"`
}
