package http

import (
	"github.com/go-playground/validator"
	"github.com/hikayatz/go-deploy-heroku/internal/app/auth"
	"github.com/hikayatz/go-deploy-heroku/internal/app/book"
	"github.com/hikayatz/go-deploy-heroku/internal/app/role"
	"github.com/hikayatz/go-deploy-heroku/internal/app/user"
	"github.com/hikayatz/go-deploy-heroku/internal/factory"
	"github.com/hikayatz/go-deploy-heroku/pkg/util"
	"github.com/labstack/echo/v4"
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
