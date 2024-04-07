package controller

import (
	"payment/internal/service"
	"payment/shared"
	"payment/shared/dto"
	"payment/shared/errors"

	"github.com/labstack/echo"
)

type (
	//PaymentController is
	PaymentController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewBeneficieryController is
func NewPaymentController(services service.Holder, deps shared.Deps) (*PaymentController, error) {
	return &PaymentController{
		services: services,
		deps:     deps,
	}, nil
}

// GetPaymentByInternalOrderID is
func (ctrl *PaymentController) GetBeneficieryAccounNumber(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyPrixaContext(ctx)
		context = pctx.Request().Context()
	)

	accountNumber := ctx.QueryParam("account_number")
	bank_id := ctx.QueryParam("bank_id")
	beneficieryAccount, err := ctrl.services.PaymentService.GetBeneficieryByAccountNumber(context, accountNumber, bank_id)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(beneficieryAccount)
}

// Transfer is
func (ctrl *PaymentController) TransferOrDisburse(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyPrixaContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.TransactionRequest)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	response, err := ctrl.services.PaymentService.InsertTransfer(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(response)
}

// Transfer is
func (ctrl *PaymentController) CallbackTransfer(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyPrixaContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.CallbackTransfer)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	err := ctrl.services.PaymentService.CallBackUpdateStatusTransfer(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}
