package shared

import (
	"payment/shared/errors"

	"github.com/labstack/echo"
)

type (
	PaymentContext struct {
		echo.Context
	}
	Success struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}

	Failed struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
)

func (sc *PaymentContext) Success(data interface{}) error {
	return sc.JSON(200, Success{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func (sc *PaymentContext) SuccessWithMeta(data, meta interface{}) error {
	return sc.JSON(200, Success{
		Code:    200,
		Message: "success",
		Data:    data,
		Meta:    meta,
	})
}

func (sc *PaymentContext) Fail(err error) error {
	var (
		ed = errors.ExtractError(err)
	)

	return sc.JSON(ed.HttpCode, Failed{
		Code:    ed.Code,
		Message: "failed",
		Error:   ed.Message,
	})
}

func NewEmptyPrixaContext(parent echo.Context) *PaymentContext {
	return &PaymentContext{parent}
}
