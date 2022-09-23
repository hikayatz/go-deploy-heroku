package auth

import (
	"context"
	"errors"

	"submission-5/internal/dto"
	"submission-5/internal/factory"
	"submission-5/internal/pkg/util"
	"submission-5/internal/repository"
	"submission-5/pkg/constant"
	pkgutil "submission-5/pkg/util"
	res "submission-5/pkg/util/response"
)

type service struct {
	EmployeeRepository repository.User
}

type Service interface {
	LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error)
	RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		EmployeeRepository: f.UserRepository,
	}
}

func (s *service) LoginByEmailAndPassword(ctx context.Context, payload *dto.ByEmailAndPasswordRequest) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse

	data, err := s.EmployeeRepository.FindByEmail(ctx, &payload.Email)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return result, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.RoleID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return result, nil
}

func (s *service) RegisterByEmailAndPassword(ctx context.Context, payload *dto.RegisterUserRequestBody) (*dto.UserWithJWTResponse, error) {
	var result *dto.UserWithJWTResponse
	isExist, err := s.EmployeeRepository.ExistByEmail(ctx, &payload.Email)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	if isExist {
		return result, res.ErrorBuilder(&res.ErrorConstant.Duplicate, errors.New("user already exists"))
	}

	hashedPassword, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	payload.Password = hashedPassword

	data, err := s.EmployeeRepository.Save(ctx, payload)
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID, data.RoleID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return result, res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	result = &dto.UserWithJWTResponse{
		UserResponse: dto.UserResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		JWT: token,
	}

	return result, nil
}
