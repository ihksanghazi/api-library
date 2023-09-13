package services

import (
	"github.com/ihksanghazi/api-library/models/web"
	"github.com/ihksanghazi/api-library/repositories"
)

type UserService interface {
	GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error)
}

type UserServiceImpl struct {
	repository *repositories.Query
}

func NewUserService(repository *repositories.Query) UserService {
	return &UserServiceImpl{
		repository: repository,
	}
}

func (u *UserServiceImpl) GetAllUserService(page int, limit int) (users []web.UsersWebResponse, totalPage int64, err error) {
	var result []web.UsersWebResponse
	//pagination
	offset := (page - 1) * limit
	// getall user by page
	user := u.repository.User
	Count, errRepository := user.Select(user.ID, user.Username, user.Email, user.PhoneNumber, user.Address, user.ImageUrl).ScanByPage(&result, offset, limit)

	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return result, TotalPage, errRepository
}
