// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package repositories

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/ihksanghazi/api-library/models/domain"
)

func newBook(db *gorm.DB, opts ...gen.DOOption) book {
	_book := book{}

	_book.bookDo.UseDB(db, opts...)
	_book.bookDo.UseModel(&domain.Book{})

	tableName := _book.bookDo.TableName()
	_book.ALL = field.NewAsterisk(tableName)
	_book.ID = field.NewField(tableName, "id")
	_book.Title = field.NewString(tableName, "title")
	_book.Author = field.NewString(tableName, "author")
	_book.PublicationYear = field.NewTime(tableName, "publication_year")
	_book.ImageUrl = field.NewString(tableName, "image_url")
	_book.Total = field.NewInt(tableName, "total")
	_book.CreatedAt = field.NewTime(tableName, "created_at")
	_book.UpdatedAt = field.NewTime(tableName, "updated_at")
	_book.DeletedAt = field.NewField(tableName, "deleted_at")

	_book.fillFieldMap()

	return _book
}

type book struct {
	bookDo

	ALL             field.Asterisk
	ID              field.Field
	Title           field.String
	Author          field.String
	PublicationYear field.Time
	ImageUrl        field.String
	Total           field.Int
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field

	fieldMap map[string]field.Expr
}

func (b book) Table(newTableName string) *book {
	b.bookDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b book) As(alias string) *book {
	b.bookDo.DO = *(b.bookDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *book) updateTableName(table string) *book {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewField(table, "id")
	b.Title = field.NewString(table, "title")
	b.Author = field.NewString(table, "author")
	b.PublicationYear = field.NewTime(table, "publication_year")
	b.ImageUrl = field.NewString(table, "image_url")
	b.Total = field.NewInt(table, "total")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")

	b.fillFieldMap()

	return b
}

func (b *book) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *book) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 9)
	b.fieldMap["id"] = b.ID
	b.fieldMap["title"] = b.Title
	b.fieldMap["author"] = b.Author
	b.fieldMap["publication_year"] = b.PublicationYear
	b.fieldMap["image_url"] = b.ImageUrl
	b.fieldMap["total"] = b.Total
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
}

func (b book) clone(db *gorm.DB) book {
	b.bookDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b book) replaceDB(db *gorm.DB) book {
	b.bookDo.ReplaceDB(db)
	return b
}

type bookDo struct{ gen.DO }

type IBookDo interface {
	gen.SubQuery
	Debug() IBookDo
	WithContext(ctx context.Context) IBookDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBookDo
	WriteDB() IBookDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBookDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBookDo
	Not(conds ...gen.Condition) IBookDo
	Or(conds ...gen.Condition) IBookDo
	Select(conds ...field.Expr) IBookDo
	Where(conds ...gen.Condition) IBookDo
	Order(conds ...field.Expr) IBookDo
	Distinct(cols ...field.Expr) IBookDo
	Omit(cols ...field.Expr) IBookDo
	Join(table schema.Tabler, on ...field.Expr) IBookDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBookDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBookDo
	Group(cols ...field.Expr) IBookDo
	Having(conds ...gen.Condition) IBookDo
	Limit(limit int) IBookDo
	Offset(offset int) IBookDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBookDo
	Unscoped() IBookDo
	Create(values ...*domain.Book) error
	CreateInBatches(values []*domain.Book, batchSize int) error
	Save(values ...*domain.Book) error
	First() (*domain.Book, error)
	Take() (*domain.Book, error)
	Last() (*domain.Book, error)
	Find() ([]*domain.Book, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Book, err error)
	FindInBatches(result *[]*domain.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*domain.Book) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBookDo
	Assign(attrs ...field.AssignExpr) IBookDo
	Joins(fields ...field.RelationField) IBookDo
	Preload(fields ...field.RelationField) IBookDo
	FirstOrInit() (*domain.Book, error)
	FirstOrCreate() (*domain.Book, error)
	FindByPage(offset int, limit int) (result []*domain.Book, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBookDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b bookDo) Debug() IBookDo {
	return b.withDO(b.DO.Debug())
}

func (b bookDo) WithContext(ctx context.Context) IBookDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b bookDo) ReadDB() IBookDo {
	return b.Clauses(dbresolver.Read)
}

func (b bookDo) WriteDB() IBookDo {
	return b.Clauses(dbresolver.Write)
}

func (b bookDo) Session(config *gorm.Session) IBookDo {
	return b.withDO(b.DO.Session(config))
}

func (b bookDo) Clauses(conds ...clause.Expression) IBookDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b bookDo) Returning(value interface{}, columns ...string) IBookDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b bookDo) Not(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b bookDo) Or(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b bookDo) Select(conds ...field.Expr) IBookDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b bookDo) Where(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b bookDo) Order(conds ...field.Expr) IBookDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b bookDo) Distinct(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b bookDo) Omit(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b bookDo) Join(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b bookDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b bookDo) RightJoin(table schema.Tabler, on ...field.Expr) IBookDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b bookDo) Group(cols ...field.Expr) IBookDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b bookDo) Having(conds ...gen.Condition) IBookDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b bookDo) Limit(limit int) IBookDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b bookDo) Offset(offset int) IBookDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b bookDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBookDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b bookDo) Unscoped() IBookDo {
	return b.withDO(b.DO.Unscoped())
}

func (b bookDo) Create(values ...*domain.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b bookDo) CreateInBatches(values []*domain.Book, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b bookDo) Save(values ...*domain.Book) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b bookDo) First() (*domain.Book, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Book), nil
	}
}

func (b bookDo) Take() (*domain.Book, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Book), nil
	}
}

func (b bookDo) Last() (*domain.Book, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Book), nil
	}
}

func (b bookDo) Find() ([]*domain.Book, error) {
	result, err := b.DO.Find()
	return result.([]*domain.Book), err
}

func (b bookDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*domain.Book, err error) {
	buf := make([]*domain.Book, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b bookDo) FindInBatches(result *[]*domain.Book, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b bookDo) Attrs(attrs ...field.AssignExpr) IBookDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b bookDo) Assign(attrs ...field.AssignExpr) IBookDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b bookDo) Joins(fields ...field.RelationField) IBookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b bookDo) Preload(fields ...field.RelationField) IBookDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b bookDo) FirstOrInit() (*domain.Book, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Book), nil
	}
}

func (b bookDo) FirstOrCreate() (*domain.Book, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*domain.Book), nil
	}
}

func (b bookDo) FindByPage(offset int, limit int) (result []*domain.Book, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b bookDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b bookDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b bookDo) Delete(models ...*domain.Book) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *bookDo) withDO(do gen.Dao) *bookDo {
	b.DO = *do.(*gen.DO)
	return b
}
