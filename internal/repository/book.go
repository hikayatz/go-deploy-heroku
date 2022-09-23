package repository

import (
	"context"
	"gorm.io/gorm"
	"strings"
	"submission-5/internal/dto"
	"submission-5/internal/model"
	pkgdto "submission-5/pkg/dto"
)

type Book interface {
	FindAll(ctx context.Context, payload *dto.SearchGetBookRequest) (*[]model.Book, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint) (*model.Book, error)
	Save(ctx context.Context, payload *dto.BookRequestBody) (*model.Book, error)
	Update(ctx context.Context, id uint, payload *dto.UpdateBookRequestBody) (*model.Book, error)
	Destroy(ctx context.Context, id uint) error
}

type book struct {
	Db *gorm.DB
}

func NewBookRepository(db *gorm.DB) Book {
	return &book{
		db,
	}
}

func (b book) FindAll(ctx context.Context, payload *dto.SearchGetBookRequest) (*[]model.Book, *pkgdto.PaginationInfo, error) {
	var books []model.Book
	var count int64

	query := b.Db.WithContext(ctx).Model(&model.Book{})
	if payload.Title != "" {
		query = query.Where("lower(title) LIKE ?  ", "%"+strings.ToLower(payload.Title)+"%")
	}
	if payload.Author != "" {
		query = query.Where("lower(author) LIKE ?  ", "%"+strings.ToLower(payload.Author)+"%")
	}
	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}
	limit, offset := pkgdto.GetLimitOffset(&payload.Pagination)

	err := query.Limit(limit).Offset(offset).Find(&books).Error

	return &books, pkgdto.CheckInfoPagination(&payload.Pagination, count), err
}

func (b book) FindByID(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	q := b.Db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", id)

	err := q.First(&book).Error
	return &book, err
}

func (b book) Save(ctx context.Context, payload *dto.BookRequestBody) (*model.Book, error) {
	newBook := model.Book{
		Title:       payload.Title,
		Author:      payload.Author,
		Description: payload.Description,
		Year:        payload.Year,
	}
	if err := b.Db.WithContext(ctx).Save(&newBook).Error; err != nil {
		return nil, err
	}
	return &newBook, nil
}

func (b book) Update(ctx context.Context, id uint, payload *dto.UpdateBookRequestBody) (*model.Book, error) {
	var book *model.Book
	q := b.Db.WithContext(ctx).Model(&model.Book{}).Where("id = ?", id)
	q2 := q
	err := q.First(&book).Error
	if err != nil {
		return nil, err
	}
	book = &model.Book{
		Title:       payload.Title,
		Author:      payload.Author,
		Description: payload.Description,
		Year:        payload.Year,
	}
	if err = q2.Updates(book).Scan(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (b book) Destroy(ctx context.Context, id uint) error {
	if err := b.Db.WithContext(ctx).Delete(&model.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}
