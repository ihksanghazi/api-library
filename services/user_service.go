package services

import (
	"context"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error)
	GetUserByIdService(id string) (user web.UserWebResponse, err error)
	UpdateUserService(id string, req web.UpdateUserWebRequest) (result web.UpdateUserWebRequest, err error)
}

type UserServiceImpl struct {
	ctx   context.Context
	db    *gorm.DB
	model domain.User
}

func NewUserService(context context.Context, db *gorm.DB, model domain.User) UserService {
	return &UserServiceImpl{
		ctx:   context,
		db:    db,
		model: model,
	}
}

func (u *UserServiceImpl) GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error) {
	var result []web.UsersWebResponse
	var Count int64
	//pagination
	offset := (page - 1) * limit
	// getall user by page
	Error := u.db.Model(u.model).WithContext(u.ctx).Count(&Count).Offset(offset).Limit(limit).Find(&result).Error

	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return result, TotalPage, Error
}

func (u *UserServiceImpl) GetUserByIdService(id string) (user web.UserWebResponse, err error) {
	var response web.UserWebResponse
	// Get User By Id
	Error := u.db.Model(u.model).WithContext(u.ctx).Where("id = ?", id).Preload("Books.Borrow").First(&response).Error
	return response, Error
}

func (u *UserServiceImpl) UpdateUserService(id string, req web.UpdateUserWebRequest) (result web.UpdateUserWebRequest, err error) {
	// parsing to domain
	u.model.Username = req.Username
	u.model.Email = req.Email
	u.model.PhoneNumber = req.PhoneNumber
	u.model.Address = req.Address
	u.model.ImageUrl = req.ImageUrl
	u.model.Password = req.Password
	if req.Password != "" {
		newPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		u.model.Password = string(newPassword)
		req.Password = string(newPassword)
	}
	// transaction
	errTransaction := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(u.model).WithContext(u.ctx).Where("id = ?", id).Updates(u.model).Error; err != nil {
			return err
		}
		return nil
	})
	return req, errTransaction
}
