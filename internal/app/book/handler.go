package book

import (
	"log"
	"net/http"

	"github.com/hikayatz/go-deploy-heroku/internal/dto"
	"github.com/hikayatz/go-deploy-heroku/internal/factory"
	"github.com/hikayatz/go-deploy-heroku/internal/pkg/enum"
	"github.com/hikayatz/go-deploy-heroku/internal/pkg/util"
	pkgdto "github.com/hikayatz/go-deploy-heroku/pkg/dto"
	res "github.com/hikayatz/go-deploy-heroku/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Get(c echo.Context) error {

	payload := new(dto.SearchGetBookRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(http.StatusOK, result.Data, "Get book success", &result.PaginationInfo).Send(c)
}

func (h *handler) GetById(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	_, err := util.ParseJWTToken(authHeader)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}

	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.FindByID(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) UpdateById(c echo.Context) error {
	payload := new(dto.UpdateBookRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	isAdminOrSameUser := (jwtClaims.UserID == *payload.ID) || (jwtClaims.RoleID == uint(enum.Admin))
	log.Println(isAdminOrSameUser)
	if (err != nil) || !isAdminOrSameUser {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}
	result, err := h.service.UpdateById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(result).Send(c)
}

func (h *handler) DeleteById(c echo.Context) error {
	payload := new(pkgdto.ByIDRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	authHeader := c.Request().Header.Get("Authorization")
	jwtClaims, err := util.ParseJWTToken(authHeader)
	if (err != nil) || (jwtClaims.UserID != payload.ID) || (jwtClaims.RoleID != uint(enum.Admin)) {
		return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
	}
	err = h.service.DeleteById(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(nil).Send(c)
}

func (h *handler) Store(c echo.Context) error {
	payload := new(dto.BookRequestBody)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	model, err := h.service.Store(c.Request().Context(), payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}
	return res.SuccessResponse(model).Send(c)
}
