package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	UpdateUserRequestBody struct {
		ID       *uint   `param:"id" validate:"required"`
		Fullname *string `json:"fullname" validate:"omitempty"`
		Email    *string `json:"email" validate:"omitempty,email"`
		Password *string `json:"password" validate:"omitempty"`
		RoleID   *uint   `json:"role_id" validate:"omitempty"`
	}
	UserResponse struct {
		ID       uint   `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
	}
	UserWithJWTResponse struct {
		UserResponse
		JWT string `json:"jwt"`
	}
	UserWithCUDResponse struct {
		UserResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
	UserDetailResponse struct {
		UserResponse
		Role RoleResponse `json:"role"`
	}
)