package employee

import (
	"latihan_service/internal/dto"
	"latihan_service/internal/middleware"
	"latihan_service/pkg/util"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.GET("", h.Get)
	g.PUT("/update/:id", h.UpdateById)

	// auth := g
	// auth.Use(mid.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	// auth.PUT("/update/:id", h.UpdateById)
}
