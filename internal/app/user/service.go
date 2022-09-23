package user

import (
	"context"

	"submission-5/internal/dto"
	"submission-5/internal/factory"
	"submission-5/internal/repository"
	"submission-5/pkg/constant"
	pkgdto "submission-5/pkg/dto"
	res "submission-5/pkg/util/response"
)

type service struct {
	EmployeeRepository repository.User
}

type Service interface {
	Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserDetailResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateUserRequestBody) (*dto.UserDetailResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserWithCUDResponse, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		EmployeeRepository: f.UserRepository,
	}
}

func (s *service) Find(ctx context.Context, payload *pkgdto.SearchGetRequest) (*pkgdto.SearchGetResponse[dto.UserResponse], error) {
	employees, info, err := s.EmployeeRepository.FindAll(ctx, payload, &payload.Pagination)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var data []dto.UserResponse

	for _, employee := range employees {
		data = append(data, dto.UserResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		})

	}

	result := new(pkgdto.SearchGetResponse[dto.UserResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserDetailResponse, error) {
	data, err := s.EmployeeRepository.FindByID(ctx, payload.ID, true)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserDetailResponse{
		UserResponse: dto.UserResponse{
			ID:       data.ID,
			Fullname: data.Fullname,
			Email:    data.Email,
		},
		Role: dto.RoleResponse{
			ID:   data.Role.ID,
			Name: data.Role.Name,
		},
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateUserRequestBody) (*dto.UserDetailResponse, error) {
	employee, err := s.EmployeeRepository.FindByID(ctx, *payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	_, err = s.EmployeeRepository.Edit(ctx, &employee, payload)
	if err != nil {
		return &dto.UserDetailResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserDetailResponse{
		UserResponse: dto.UserResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		},
		Role: dto.RoleResponse{
			ID:   employee.Role.ID,
			Name: employee.Role.Name,
		},
	}

	return result, nil
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.UserWithCUDResponse, error) {
	employee, err := s.EmployeeRepository.FindByID(ctx, payload.ID, false)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	_, err = s.EmployeeRepository.Destroy(ctx, &employee)
	if err != nil {
		return &dto.UserWithCUDResponse{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.UserWithCUDResponse{
		UserResponse: dto.UserResponse{
			ID:       employee.ID,
			Fullname: employee.Fullname,
			Email:    employee.Email,
		},
		CreatedAt: employee.CreatedAt,
		UpdatedAt: employee.UpdatedAt,
		DeletedAt: employee.DeletedAt,
	}

	return result, nil
}
