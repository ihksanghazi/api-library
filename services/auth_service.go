package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/repositories"
	"github.com/ihksanghazi/api-library/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterService(req *web.RegisterWebRequest) (*web.RegisterWebRequest, error)
	LoginService(req *web.LoginWebRequest, timeRefreshToken *time.Time, timeAccessToken *time.Time) (accessToken *string, refreshToken *string, err error)
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

func (a *AuthServiceImpl) RegisterService(req *web.RegisterWebRequest) (*web.RegisterWebRequest, error) {
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
		req.Role = "user"
	} else {
		user.Role = req.Role
	}

	// Transaction
	errTrans := a.repository.Transaction(func(tx *repositories.Query) error {
		// find user by email if already return error
		_, errRep := tx.User.WithContext(a.ctx).Where(tx.User.Email.Eq(req.Email)).First()
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

		user.Password = string(hashPassword)
		req.Password = string(hashPassword)

		// insert user to database
		if errCreate := tx.User.WithContext(a.ctx).Create(&user); errCreate != nil {
			return errCreate
		}

		return nil
	})

	return req, errTrans
}

func (a *AuthServiceImpl) LoginService(req *web.LoginWebRequest, timeRefreshToken *time.Time, timeAccessToken *time.Time) (accessToken *string, refreshToken *string, err error) {
	var AccessToken, RefreshToken string
	//transaction
	errTransaction := a.repository.Transaction(func(tx *repositories.Query) error {
		// find user by email
		user, errQuery := tx.User.WithContext(a.ctx).Where(tx.User.Email.Eq(req.Email)).First()
		if errQuery == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		} else if errQuery != nil {
			return errQuery
		}

		// matching password
		if errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); errHash != nil {
			return errors.New("wrong password")
		}

		// // Generate Refresh Token
		refreshToken, errRefreshToken := utils.GenerateToken(user, os.Getenv("REFRESH_TOKEN_JWT"), *timeRefreshToken)
		if errRefreshToken != nil {
			return errRefreshToken
		}

		//insert refresh token to current user
		_, errUpdate := tx.User.WithContext(a.ctx).Where(tx.User.ID.Eq(user.ID)).Update(tx.User.RefreshToken, refreshToken)
		if errUpdate != nil {
			return errUpdate
		}

		RefreshToken = refreshToken

		// generate access Token
		accessToken, errAccessToken := utils.GenerateToken(user, os.Getenv("ACCESS_TOKEN_JWT"), *timeAccessToken)
		if errAccessToken != nil {
			return errAccessToken
		}

		AccessToken = accessToken
		return nil
	})

	return &AccessToken, &RefreshToken, errTransaction
}
