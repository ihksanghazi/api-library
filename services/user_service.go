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
	DeleteUserService(id string) (err error)
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
	var domain domain.User
	domain.Username = req.Username
	domain.Email = req.Email
	domain.PhoneNumber = req.PhoneNumber
	domain.Address = req.Address
	domain.ImageUrl = req.ImageUrl
	domain.Password = req.Password
	if req.Password != "" {
		newPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		domain.Password = string(newPassword)
		req.Password = string(newPassword)
	}
	// transaction
	errTransaction := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(u.model).WithContext(u.ctx).Where("id = ?", id).Updates(domain).Error; err != nil {
			return err
		}
		return nil
	})
	return req, errTransaction
}

func (u *UserServiceImpl) DeleteUserService(id string) (err error) {
	// transaction
	errTransaction := u.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(u.model).WithContext(u.ctx).Where("id = ?", id).Delete(&u.model).Error; err != nil {
			return err
		}
		return nil
	})
	return errTransaction
}
