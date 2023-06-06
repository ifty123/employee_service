package middleware

import (
	"latihan_service/internal/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Middleware struct {
}

func NewMidleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,PATCH")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Authorization,Origin,Accept,datetime,signature,Content-Type")
		c.Response().Header().Set("Content-Type", "application/json")
		return next(c)
	}
}

func (m *Middleware) JWTMiddleware(claims dto.JWTClaims, signingKey []byte) echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &dto.JWTClaims{},
		SigningKey: signingKey,
	}
	return middleware.JWTWithConfig(config)
}

func (m *Middleware) LogMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_rfc3339}] ${status} ${method} ${host}${uri} ${latency_human}` + "\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
		// Output:           f,
	}))
}
