package controller

import (
	"net/http"
	"payment/shared"
	"payment/shared/errors"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		Deps shared.Deps

		PaymentController *PaymentController
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewPaymentController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {

	app.Validator = h.Deps.CustomValidator
	app.HTTPErrorHandler = h.ErrorHandler

	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	v1 := app.Group("/v1")

	paymentRoute := v1.Group("/payment")
	paymentRoute.POST("/transfer-disbursement", h.PaymentController.TransferOrDisburse)
	paymentRoute.GET("/", h.PaymentController.GetBeneficieryAccounNumber)
	paymentRoute.PUT("/callback-transfer-disbursement", h.PaymentController.CallbackTransfer)

}

func (h *Holder) ErrorHandler(err error, ctx echo.Context) {
	var (
		sctx, ok = ctx.(*shared.PaymentContext)
	)

	if !ok {
		sctx = shared.NewEmptyPrixaContext(ctx)
	}

	e, ok := err.(*echo.HTTPError)
	if ok {
		msg, ok := e.Message.(string)
		if !ok {
			msg = err.Error()
		}
		err = errors.ErrBase.New(msg).WithProperty(errors.ErrCodeProperty, e.Code).WithProperty(errors.ErrHttpCodeProperty, e.Code)
	}

	h.Deps.Logger.Errorf(
		"path=%s error=%s",
		sctx.Path(),
		err,
	)

	_ = sctx.Fail(err)
}
