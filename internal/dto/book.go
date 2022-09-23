package dto

import "submission-5/pkg/dto"

type (
	SearchGetBookRequest struct {
		dto.Pagination
		Title    string   `query:"title"`
		Author   string   `query:"author"`
		AscField []string `query:"asc_field"`
		DscField []string `query:"dsc_field"`
	}
	BookRequestBody struct {
		Title       string `json:"title" validate:"required"`
		Author      string `json:"author" validate:"required"`
		Description string `json:"description" validate:"required"`
		Year        int    `json:"year"`
	}
	BookResponse struct {
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Author      string `json:"author"`
		Description string `json:"description"`
		Year        int    `json:"year"`
	}
	UpdateBookRequestBody struct {
		ID *uint `param:"id" validate:"required"`
		BookRequestBody
	}
)
