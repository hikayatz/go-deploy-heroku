package book

import (
	"context"
	"submission-5/internal/dto"
	"submission-5/internal/factory"
	"submission-5/internal/repository"
	"submission-5/pkg/constant"
	pkgdto "submission-5/pkg/dto"
	res "submission-5/pkg/util/response"
)

type Service interface {
	Find(ctx context.Context, payload *dto.SearchGetBookRequest) (*pkgdto.SearchGetResponse[dto.BookResponse], error)
	FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookResponse, error)
	UpdateById(ctx context.Context, payload *dto.UpdateBookRequestBody) (*dto.BookResponse, error)
	DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) error
	Store(ctx context.Context, payload *dto.BookRequestBody) (*dto.BookResponse, error)
}
type service struct {
	BookRepository repository.Book
}

func NewService(f *factory.Factory) Service {
	return &service{
		f.BookRepository,
	}
}
func (s *service) Find(ctx context.Context, payload *dto.SearchGetBookRequest) (*pkgdto.SearchGetResponse[dto.BookResponse], error) {
	books, info, err := s.BookRepository.FindAll(ctx, payload)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	var data []dto.BookResponse

	for _, item := range *books {
		data = append(data, dto.BookResponse{
			ID:          item.ID,
			Title:       item.Title,
			Author:      item.Author,
			Description: item.Description,
		})

	}
	result := new(pkgdto.SearchGetResponse[dto.BookResponse])
	result.Data = data
	result.PaginationInfo = *info

	return result, nil
}

func (s *service) FindByID(ctx context.Context, payload *pkgdto.ByIDRequest) (*dto.BookResponse, error) {
	data, err := s.BookRepository.FindByID(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.BookResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Year:        data.Year,
	}

	return result, nil
}

func (s *service) UpdateById(ctx context.Context, payload *dto.UpdateBookRequestBody) (*dto.BookResponse, error) {

	data, err := s.BookRepository.Update(ctx, *payload.ID, payload)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return nil, res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result := &dto.BookResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Year:        data.Year,
	}

	return result, nil
}

func (s *service) DeleteById(ctx context.Context, payload *pkgdto.ByIDRequest) error {

	err := s.BookRepository.Destroy(ctx, payload.ID)
	if err != nil {
		if err == constant.RECORD_NOT_FOUND {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	return nil
}

func (s *service) Store(ctx context.Context, payload *dto.BookRequestBody) (*dto.BookResponse, error) {
	model, err := s.BookRepository.Save(ctx, payload)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)

	}
	result := &dto.BookResponse{
		ID:          model.ID,
		Title:       model.Title,
		Author:      model.Author,
		Description: model.Description,
		Year:        model.Year,
	}
	return result, nil

}