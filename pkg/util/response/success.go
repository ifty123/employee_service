package response

import (
	"latihan_service/pkg/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

type successConstant struct {
	OK Success
}

type successResponse struct {
	Success bool                `json:"success" default:"true"`
	Message string              `json:"message" default:"true"`
	Meta    *dto.PaginationInfo `json:"meta"`
	Data    interface{}         `json:"data"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

var SuccessConstant successConstant = successConstant{
	OK: Success{
		Response: successResponse{
			Success: true,
			Message: "Request successfully proceed",
			Data:    nil,
		},
		Code: http.StatusOK,
	},
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string, info *dto.PaginationInfo) *Success {
	return &Success{
		Response: successResponse{
			Success: true,
			Message: message,
			Meta:    info,
			Data:    data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant.OK, data)
}

func (s *Success) Send(c echo.Context) error {
	return c.JSON(s.Code, s.Response)
}
