package services

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	RegisterService(req *web.RegisterWebRequest) (*web.RegisterWebRequest, error)
	LoginService(req *web.LoginWebRequest, timeRefreshToken *time.Time, timeAccessToken *time.Time) (accessToken *string, refreshToken *string, err error)
	GetTokenService(refreshToken *string) (accessToken *string, err error)
}

type AuthServiceImpl struct {
	ctx   context.Context
	db    *gorm.DB
	model domain.User
}

func NewAuthService(ctx context.Context, db *gorm.DB, model domain.User) AuthService {
	return &AuthServiceImpl{
		ctx:   ctx,
		db:    db,
		model: model,
	}
}

func (a *AuthServiceImpl) RegisterService(req *web.RegisterWebRequest) (*web.RegisterWebRequest, error) {
	// passing data from request to domain
	var model domain.User
	model.Username = req.Username
	model.Email = req.Email
	model.Password = req.Password
	model.Address = req.Address
	model.PhoneNumber = req.PhoneNumber
	model.ImageUrl = req.ImageUrl
	if req.Role == "" {
		model.Role = "user"
		req.Role = "user"
	} else {
		model.Role = req.Role
	}

	// Transaction
	errTrans := a.db.Transaction(func(tx *gorm.DB) error {
		// find user by email if already return error
		if errRep := tx.Model(a.model).WithContext(a.ctx).Where("email = ?", req.Email).First(&model).Error; errRep == nil {
			return errors.New("email is already used")
		} else if errRep != nil && errRep != gorm.ErrRecordNotFound {
			return errRep
		}

		// hashing password
		hashPassword, errHashPass := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if errHashPass != nil {
			return errHashPass
		}

		model.Password = string(hashPassword)
		req.Password = string(hashPassword)
		// insert user to database
		if errCreate := tx.Model(a.model).WithContext(a.ctx).Create(&model).Error; errCreate != nil {
			return errCreate
		}

		return nil
	})

	return req, errTrans
}

func (a *AuthServiceImpl) LoginService(req *web.LoginWebRequest, timeRefreshToken *time.Time, timeAccessToken *time.Time) (accessToken *string, refreshToken *string, err error) {
	var AccessToken, RefreshToken string
	//transaction
	errTransaction := a.db.Transaction(func(tx *gorm.DB) error {
		// find user by email
		errQuery := tx.Model(a.model).WithContext(a.ctx).Where("email = ?", req.Email).First(&a.model).Error
		if errQuery == gorm.ErrRecordNotFound {
			return errors.New("user not found")
		} else if errQuery != nil {
			return errQuery
		}

		// matching password
		if errHash := bcrypt.CompareHashAndPassword([]byte(a.model.Password), []byte(req.Password)); errHash != nil {
			return errors.New("wrong password")
		}

		// // Generate Refresh Token
		refreshToken, errRefreshToken := utils.GenerateToken(a.model.ID.String(), a.model.Username, a.model.Email, os.Getenv("REFRESH_TOKEN_JWT"), *timeRefreshToken)
		if errRefreshToken != nil {
			return errRefreshToken
		}

		//insert refresh token to current user
		errUpdate := tx.Model(a.model).WithContext(a.ctx).Where("id = ?", a.model.ID.String()).Update("refresh_token", refreshToken).Error
		if errUpdate != nil {
			return errUpdate
		}

		RefreshToken = refreshToken

		// generate access Token
		accessToken, errAccessToken := utils.GenerateToken(a.model.ID.String(), a.model.Username, a.model.Email, os.Getenv("ACCESS_TOKEN_JWT"), *timeAccessToken)
		if errAccessToken != nil {
			return errAccessToken
		}

		AccessToken = accessToken
		return nil
	})

	return &AccessToken, &RefreshToken, errTransaction
}

func (a *AuthServiceImpl) GetTokenService(refreshToken *string) (accessToken *string, err error) {
	// parsing refresh token
	RefreshToken, errRefreshToken := utils.ParsingToken(*refreshToken, os.Getenv("REFRESH_TOKEN_JWT"))
	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	// generate access token
	AccessToken, errAccessToken := utils.GenerateToken(RefreshToken.ID, RefreshToken.Username, RefreshToken.Email, os.Getenv("ACCESS_TOKEN_JWT"), time.Now().Add(time.Second*20))
	if errAccessToken != nil {
		return nil, errAccessToken
	}

	return &AccessToken, nil
}
