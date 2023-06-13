package employee

import (
	"errors"
	"latihan_service/internal/dto"
	"latihan_service/internal/factory"
	pkgdto "latihan_service/pkg/dto"
	"latihan_service/pkg/util"
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

func (h *handler) UpdateById(c echo.Context) error {
	payload := new(dto.UpdateEmployeeReq)

	if err := c.Bind(payload); err != nil {
		log.Println("err bind :", err)
		return response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		log.Println("error :", err)
		return response.ErrorBuilder(&response.ErrorConstant.Validation, err).Send(c)
	}

	//auth
	authHeader := c.Request().Header.Get("Authorization")
	log.Println("isi jwt :", authHeader)

	jwtClaims, err := util.ParseJWTToken(authHeader)
	if err != nil {
		log.Println("-> err parse http :", err)
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, err).Send(c)
	}

	log.Println("isi jwt claim :", jwtClaims.ID, "-", jwtClaims.Email)

	if payload.ID != jwtClaims.UserID {
		return response.ErrorBuilder(&response.ErrorConstant.Unauthorized, errors.New("Id not match")).Send(c)
	}

	payload.ID = jwtClaims.UserID

	res, err := h.service.UpdateById(c.Request().Context(), payload)

	if err != nil {
		log.Println("=> err service :", err)
		return response.ErrorResponse(err).Send(c)
	}

	return response.SuccessResponse(res).Send(c)
}
