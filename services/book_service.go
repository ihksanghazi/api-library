package services

import (
	"context"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBookService(req web.CreateBookWebRequest) (result web.CreateBookWebRequest, err error)
	GetAllBookService(page int, limit int) (result []web.GetAllBooksWebResponse, totalPage int64, err error)
	UpdateBookService(id string, req web.UpdateBookWebRequest) (result web.UpdateBookWebRequest, err error)
}

type BookServiceImpl struct {
	db    *gorm.DB
	ctx   context.Context
	model domain.Book
}

func NewBookService(db *gorm.DB, ctx context.Context, model domain.Book) BookService {
	return &BookServiceImpl{
		db:    db,
		ctx:   ctx,
		model: model,
	}
}

func (b *BookServiceImpl) CreateBookService(req web.CreateBookWebRequest) (result web.CreateBookWebRequest, err error) {
	// parsing to model
	var model domain.Book
	model.Title = req.Title
	model.Author = req.Author
	model.PublicationYear = req.PublicationYear
	model.Total = req.Total
	model.ImageUrl = req.ImageUrl
	// transaction
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(model).WithContext(b.ctx).Create(&model).Error; err != nil {
			return err
		}
		return nil
	})
	return req, errTransaction
}

func (b *BookServiceImpl) GetAllBookService(page int, limit int) (result []web.GetAllBooksWebResponse, totalPage int64, err error) {
	var response []web.GetAllBooksWebResponse
	var Count int64
	//pagination
	offset := (page - 1) * limit
	// getall user by page
	error := b.db.Model(b.model).WithContext(b.ctx).Count(&Count).Offset(offset).Limit(limit).Find(&response).Error

	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return response, TotalPage, error
}

func (b *BookServiceImpl) UpdateBookService(id string, req web.UpdateBookWebRequest) (result web.UpdateBookWebRequest, err error) {
	// parsing to model
	var model domain.Book
	model.Title = req.Title
	model.Author = req.Author
	model.PublicationYear = req.PublicationYear
	model.Total = req.Total
	model.ImageUrl = req.ImageUrl
	// transaction
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(b.model).WithContext(b.ctx).Where("id = ?", id).Updates(model).Error; err != nil {
			return err
		}
		return nil
	})
	return req, errTransaction
}
