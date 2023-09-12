package web

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
