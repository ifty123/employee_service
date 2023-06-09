package http

import (
	"latihan_service/internal/app/auth"
	"latihan_service/internal/app/employee"
	"latihan_service/internal/factory"
	"latihan_service/pkg/util"

	"github.com/go-playground/validator"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {

	e.Validator = &util.CustomValidator{Validator: validator.New()}
	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	//group
	v1 := e.Group("/api/v1")
	employee.NewHandler(f).Route(v1.Group("/employees"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
}
