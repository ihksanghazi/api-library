package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ihksanghazi/api-library/models/domain"
	"github.com/ihksanghazi/api-library/models/web"
	"gorm.io/gorm"
)

type BookService interface {
	CreateBookService(req web.CreateBookWebRequest) (result web.CreateBookWebRequest, err error)
	GetAllBookService(page int, limit int) (result []web.BooksWebResponse, totalPage int64, err error)
	UpdateBookService(id string, req web.UpdateBookWebRequest) (result web.UpdateBookWebRequest, err error)
	DeleteBookService(id string) (err error)
	GetBookByIdService(id string) (result web.BookWebResponse, err error)
	BorrowBookService(userId string, bookId string) (err error)
	ReturnBookService(bookId string, userId string) (err error)
	GetAllExpiredService(page int, limit int) (result []web.BorrowsWebResponse, totalPage int64, err error)
	UpdateExpiredService() (err error)
}

type BookServiceImpl struct {
	db     *gorm.DB
	ctx    context.Context
	book   domain.Book
	borrow domain.Borrowing
}

func NewBookService(db *gorm.DB, ctx context.Context, book domain.Book, borrow domain.Borrowing) BookService {
	return &BookServiceImpl{
		db:     db,
		ctx:    ctx,
		book:   book,
		borrow: borrow,
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

func (b *BookServiceImpl) GetAllBookService(page int, limit int) (result []web.BooksWebResponse, totalPage int64, err error) {
	var response []web.BooksWebResponse
	var Count int64
	//pagination
	offset := (page - 1) * limit
	// getall user by page
	error := b.db.Model(b.book).WithContext(b.ctx).Count(&Count).Offset(offset).Limit(limit).Find(&response).Error

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
		if err := tx.Model(b.book).WithContext(b.ctx).Where("id = ?", id).Updates(model).Error; err != nil {
			return err
		}
		return nil
	})
	return req, errTransaction
}

func (b *BookServiceImpl) DeleteBookService(id string) (err error) {
	// transaction
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(b.book).WithContext(b.ctx).Where("id = ?", id).Delete(&b.book).Error; err != nil {
			return err
		}
		return nil
	})
	return errTransaction
}

func (b *BookServiceImpl) GetBookByIdService(id string) (result web.BookWebResponse, err error) {
	var model web.BookWebResponse
	Error := b.db.Model(b.book).WithContext(b.ctx).Where("id = ?", id).Preload("Users.Borrow").First(&model).Error
	return model, Error
}

func (b *BookServiceImpl) BorrowBookService(userId string, bookId string) (err error) {
	UserId := uuid.MustParse(userId)
	BookId := uuid.MustParse(bookId)

	borrowingDate := time.Now()
	returnDate := time.Now().Add(time.Hour * 24 * 7)

	var book domain.Book

	// parsing to model
	var borrow domain.Borrowing
	borrow.BookID = BookId
	borrow.UserID = UserId
	borrow.BorrowingDate = borrowingDate
	borrow.ReturnDate = returnDate
	borrow.Status = "borrowed"
	// transaction
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if errSelect := tx.Model(b.book).WithContext(b.ctx).Where("id = ?", bookId).First(&book).Error; errSelect != nil {
			return errSelect
		}
		if book.Total == 0 {
			return errors.New("book is not available for borrowing")
		}
		if errCreate := tx.Model(b.borrow).WithContext(b.ctx).Create(&borrow).Error; errCreate != nil {
			return errCreate
		}
		if errUpdate := tx.Model(b.book).WithContext(b.ctx).Where("id = ?", bookId).Update("total", gorm.Expr("total - ?", 1)).Error; errUpdate != nil {
			return errUpdate
		}
		return nil
	})

	return errTransaction
}

func (b *BookServiceImpl) ReturnBookService(bookId string, userId string) (err error) {
	//transaction
	var borrowing domain.Borrowing
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if errFind := tx.Model(b.borrow).WithContext(b.ctx).Where("user_id = ? AND book_id = ?", userId, bookId).First(&borrowing).Error; errFind != nil {
			return errFind
		}
		if errDelete := tx.Model(b.borrow).WithContext(b.ctx).Where("user_id = ? AND book_id = ?", userId, bookId).Delete(&borrowing).Error; errDelete != nil {
			return errDelete
		}
		if errUpdate := tx.Model(b.book).WithContext(b.ctx).Where("id = ?", bookId).Update("total", gorm.Expr("total + ?", 1)).Error; errUpdate != nil {
			return errUpdate
		}
		return nil
	})

	return errTransaction
}

func (b *BookServiceImpl) GetAllExpiredService(page int, limit int) (result []web.BorrowsWebResponse, totalPage int64, err error) {
	// Get All Borrow
	var response []web.BorrowsWebResponse

	var Count int64
	//pagination
	offset := (page - 1) * limit
	Error := b.db.Model(b.borrow).WithContext(b.ctx).Where("status = ?", "expired").Preload("User").Preload("Book").Count(&Count).Offset(offset).Limit(limit).Find(&response).Error
	TotalPage := (Count + int64(limit) - 1) / int64(limit)

	return response, TotalPage, Error
}

func (b *BookServiceImpl) UpdateExpiredService() (err error) {
	errTransaction := b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(b.ctx).Exec("update borrowings set status = 'expired' where age(now(),return_date) > interval '0 days'").Error; err != nil {
			return err
		}
		return nil
	})

	return errTransaction
}
