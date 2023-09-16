package services

import (
	"context"

	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"gorm.io/gorm"
)

type BorrowService interface {
	GetBorrowsService(page int, limit int) (result []web.BorrowsWebResponse, totalPage int64, err error)
}

type BorrowServiceImpl struct {
	db    *gorm.DB
	ctx   context.Context
	model domain.Borrowing
}

func NewBorrowService(db *gorm.DB, ctx context.Context, model domain.Borrowing) BorrowService {
	return &BorrowServiceImpl{
		db:    db,
		ctx:   ctx,
		model: model,
	}
}

func (b *BorrowServiceImpl) GetBorrowsService(page int, limit int) (result []web.BorrowsWebResponse, totalPage int64, err error) {
	// Get All Borrow
	var response []web.BorrowsWebResponse

	var Count int64
	//pagination
	offset := (page - 1) * limit
	Error := b.db.Model(b.model).WithContext(b.ctx).Preload("User").Preload("Book").Count(&Count).Offset(offset).Limit(limit).Find(&response).Error
	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return response, TotalPage, Error
}
