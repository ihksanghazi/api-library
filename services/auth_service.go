package services

import (
	"context"
	"errors"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/repositories"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterService(req web.RegisterWebRequest) (*domain.User, error)
}

type AuthServiceImpl struct {
	ctx context.Context
}

func NewAuthService(ctx context.Context) AuthService {
	return &AuthServiceImpl{
		ctx: ctx,
	}
}

func (a *AuthServiceImpl) RegisterService(req web.RegisterWebRequest) (*domain.User, error) {
	// find user by email if already return error
	_, errRep := repositories.User.WithContext(a.ctx).Where(repositories.User.Email.Eq(req.Email)).First()
	if errRep == nil {
		return nil, errors.New("email is already used")
	} else if errRep != nil && errRep != gorm.ErrRecordNotFound {
		return nil, errRep
	}

	// hashing password
	hashPassword, errHashPass := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if errHashPass != nil {
		return nil, errHashPass
	}

	req.Password = string(hashPassword)

	// create user
	var user domain.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password
	user.Address = req.Address
	user.PhoneNumber = req.PhoneNumber
	user.ImageUrl = req.ImageUrl
	if req.Role == "" {
		user.Role = "user"
	} else {
		user.Role = req.Role
	}

	errCreate := repositories.User.WithContext(a.ctx).Create(&user)
	if errCreate != nil {
		return nil, errCreate
	}

	return &user, nil
}
