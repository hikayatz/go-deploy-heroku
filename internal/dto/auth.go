package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/golang-jwt/jwt/v4"
)

type (
	RegisterUserRequestBody struct {
		Fullname string `json:"fullname" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		RoleID   *uint  `json:"role_id"`
	}

	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		UserID uint   `json:"user_id"`
		Email  string `json:"email"`
		RoleID uint   `json:"role_id"`
		jwt.RegisteredClaims
	}
)

func (r *RegisterUserRequestBody) FillDefaults() {
	var defaultRoleID uint = 1
	if r.RoleID == nil {
		r.RoleID = &defaultRoleID
	}
}

func (req ByEmailAndPasswordRequest) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Email,
			validation.Required,
			is.Email,
		),
		validation.Field(&req.Password,
			validation.Required,
		),

	)
}