package auth

import (
	"latihan_service/internal/dto"
	"latihan_service/internal/factory"
	"latihan_service/pkg/util/response"
	"log"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service AuthService
}

func NewHandler(f *factory.Factory) *Handler {
	return &Handler{
		service: NewService(f),
	}
}

func (h *Handler) LoginByEmailAndPassword(c echo.Context) error {
	payload := new(dto.EmailAndPasswordReq)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	employee, err := h.service.LoginByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		log.Println("--> error :", err)
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(employee).Send(c)
}

func (h *Handler) RegisterByEmailAndPassword(c echo.Context) error {
	payload := new(dto.RegisterEmployeeReq)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	employee, err := h.service.RegisterByEmailAndPassword(c.Request().Context(), payload)
	if err != nil {
		log.Println("--> error :", err)
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(employee).Send(c)
}
