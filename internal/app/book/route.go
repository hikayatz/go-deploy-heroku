package book

import (
	"github.com/labstack/echo/v4"
	"submission-5/internal/dto"
	"submission-5/internal/middleware"
	"submission-5/internal/pkg/util"
)

func (h *handler) Route(g *echo.Group) {
	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.GET("", h.Get)
	g.GET("/:id", h.GetById)
	g.PUT("/:id", h.UpdateById)
	g.DELETE("/:id", h.DeleteById)
	g.POST("", h.Store)

}