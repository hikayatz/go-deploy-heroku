package response

import (
	"github.com/hikayatz/go-deploy-heroku/pkg/dto"
)

type Meta struct {
	Success         bool                `json:"success" default:"true"`
	Message         string              `json:"message" default:"true"`
	Info            *dto.PaginationInfo `json:"info"`
	ErrorValidation interface{}         `json:"error_validation,omitempty"`
}
