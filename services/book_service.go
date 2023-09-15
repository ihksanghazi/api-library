package services

import (
	"context"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBookService(req web.CreateBookWebRequest) (result web.CreateBookWebRequest, err error)
}

type BookServiceImpl struct {
	db  *gorm.DB
	ctx context.Context
}

func NewBookService(db *gorm.DB, ctx context.Context) BookService {
	return &BookServiceImpl{
		db:  db,
		ctx: ctx,
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
