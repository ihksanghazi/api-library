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
	ctx        context.Context
	repository *repositories.Query
}

func NewAuthService(ctx context.Context, repository *repositories.Query) AuthService {
	return &AuthServiceImpl{
		ctx:        ctx,
		repository: repository,
	}
}

func (a *AuthServiceImpl) RegisterService(req web.RegisterWebRequest) (*domain.User, error) {
	// passing data from request to domain
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

	// Transaction
	errTrans := a.repository.Transaction(func(tx *repositories.Query) error {
		// find user by email if already return error
		_, errRep := tx.User.WithContext(a.ctx).Where(repositories.User.Email.Eq(req.Email)).First()
		if errRep == nil {
			return errors.New("email is already used")
		} else if errRep != nil && errRep != gorm.ErrRecordNotFound {
			return errRep
		}

		// hashing password
		hashPassword, errHashPass := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if errHashPass != nil {
			return errHashPass
		}

		req.Password = string(hashPassword)

		// insert user to database
		if errCreate := tx.User.WithContext(a.ctx).Create(&user); errCreate != nil {
			return errCreate
		}

		return nil
	})

	return &user, errTrans
}
