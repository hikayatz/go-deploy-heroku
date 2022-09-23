package gorm_scope

import (
	"gorm.io/gorm"
	"submission-5/pkg/dto"
)

func Paginate(p dto.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var (
			limit, offset int
		)
		if p.PageSize != nil {
			limit = *p.PageSize
		} else {
			limit = 10
			p.PageSize = &limit
		}

		if p.Page != nil {
			offset = (*p.Page - 1) * limit
		} else {
			offset = 0
		}

		return db.Offset(offset).Limit(limit)
	}
}
