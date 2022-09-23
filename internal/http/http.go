package http

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"submission-5/internal/app/auth"
	"submission-5/internal/app/book"
	"submission-5/internal/app/role"
	"submission-5/internal/app/user"
	"submission-5/internal/factory"
	"submission-5/pkg/util"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
	role.NewHandler(f).Route(v1.Group("/roles"))
	book.NewHandler(f).Route(v1.Group("/books"))

}
