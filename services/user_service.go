package services

import (
	"context"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error)
	GetUserByIdService(id string) (user web.UserWebResponse, err error)
}

type UserServiceImpl struct {
	repository *repositories.Query
	ctx        context.Context
	db         *gorm.DB
	model      domain.User
}

func NewUserService(repository *repositories.Query, context context.Context, db *gorm.DB, model domain.User) UserService {
	return &UserServiceImpl{
		repository: repository,
		ctx:        context,
		db:         db,
		model:      model,
	}
}

func (u *UserServiceImpl) GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error) {
	var result []web.UsersWebResponse
	//pagination
	offset := (page - 1) * limit
	// getall user by page
	Count, errRepository := u.repository.User.WithContext(u.ctx).ScanByPage(&result, offset, limit)

	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return result, TotalPage, errRepository
}

func (u *UserServiceImpl) GetUserByIdService(id string) (user web.UserWebResponse, err error) {
	var response web.UserWebResponse
	// Get User By Id
	Error := u.db.Model(u.model).WithContext(u.ctx).Where("id = ?", id).Preload("Books.Borrow").First(&response).Error
	return response, Error
}
