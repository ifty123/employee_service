package employee

import (
	"latihan_service/internal/factory"
	pkgdto "latihan_service/pkg/dto"
	"latihan_service/pkg/util/response"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: Newservice(f),
	}
}

func (h *handler) Get(c echo.Context) error {
	payload := new(pkgdto.SearchGetRequest)
	if err := c.Bind(payload); err != nil {
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		log.Println("error :", err)
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	res, err := h.service.Find(c.Request().Context(), payload)
	if err != nil {
		return response.ErrorResponse(err).Send(c)
	}

	return response.CustomSuccessBuilder(http.StatusOK, res.Data, "Get employees success", &res.PaginationInfo).Send(c)
}
