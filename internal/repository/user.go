package repository

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"submission-5/internal/dto"
	"submission-5/internal/model"
	pkgdto "submission-5/pkg/dto"
	"submission-5/pkg/util"
)

type User interface {
	FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, p *pkgdto.Pagination) ([]model.User, *pkgdto.PaginationInfo, error)
	FindByID(ctx context.Context, id uint, usePreload bool) (model.User, error)
	FindByEmail(ctx context.Context, email *string) (*model.User, error)
	ExistByEmail(ctx context.Context, email *string) (bool, error)
	ExistByID(ctx context.Context, id uint) (bool, error)
	Save(ctx context.Context, employee *dto.RegisterUserRequestBody) (model.User, error)
	Edit(ctx context.Context, oldEmployee *model.User, updateData *dto.UpdateUserRequestBody) (*model.User, error)
	Destroy(ctx context.Context, employee *model.User) (*model.User, error)
}

type user struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) User {
	return &user{
		db,
	}
}

func (r *user) FindAll(ctx context.Context, payload *pkgdto.SearchGetRequest, pagination *pkgdto.Pagination) ([]model.User, *pkgdto.PaginationInfo, error) {
	var users []model.User
	var count int64

	query := r.Db.WithContext(ctx).Model(&model.User{})

	if payload.Search != "" {
		search := "%" + strings.ToLower(payload.Search) + "%"
		query = query.Where("lower(fullname) LIKE ? or lower(email) Like ? ", search, search)
	}

	countQuery := query
	if err := countQuery.Count(&count).Error; err != nil {
		return nil, nil, err
	}

	limit, offset := pkgdto.GetLimitOffset(pagination)

	err := query.Limit(limit).Offset(offset).Find(&users).Error

	return users, pkgdto.CheckInfoPagination(pagination, count), err
}

func (r *user) FindByID(ctx context.Context, id uint, usePreload bool) (model.User, error) {
	var user model.User
	q := r.Db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id)
	if usePreload {
		q = q.Preload("Division").Preload("Role")
	}
	err := q.First(&user).Error
	return user, err
}

func (r *user) FindByEmail(ctx context.Context, email *string) (*model.User, error) {
	var data model.User
	err := r.Db.WithContext(ctx).Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) ExistByEmail(ctx context.Context, email *string) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *user) ExistByID(ctx context.Context, id uint) (bool, error) {
	var (
		count   int64
		isExist bool
	)
	if err := r.Db.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return isExist, err
	}
	if count > 0 {
		isExist = true
	}
	return isExist, nil
}

func (r *user) Save(ctx context.Context, usr *dto.RegisterUserRequestBody) (model.User, error) {
	newEmployee := model.User{
		Fullname: usr.Fullname,
		Email:    usr.Email,
		Password: usr.Password,
		RoleID:   *usr.RoleID,
	}
	if err := r.Db.WithContext(ctx).Save(&newEmployee).Error; err != nil {
		return newEmployee, err
	}
	return newEmployee, nil
}

func (r *user) Edit(ctx context.Context, oldUser *model.User, updateData *dto.UpdateUserRequestBody) (*model.User, error) {
	if updateData.Fullname != nil {
		oldUser.Fullname = *updateData.Fullname
	}
	if updateData.Email != nil {
		oldUser.Email = *updateData.Email
	}
	if updateData.Password != nil {
		hashedPassword, err := util.HashPassword(*updateData.Password)
		if err != nil {
			return nil, err
		}
		oldUser.Password = hashedPassword
	}

	if updateData.RoleID != nil {
		oldUser.RoleID = *updateData.RoleID
	}

	if err := r.Db.
		WithContext(ctx).
		Save(oldUser).
		Preload("Division").
		Preload("Role").
		Find(oldUser).
		Error; err != nil {
		return nil, err
	}

	return oldUser, nil
}

func (r *user) Destroy(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.Db.WithContext(ctx).Delete(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
