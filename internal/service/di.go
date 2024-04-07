package service

import "go.uber.org/dig"

type (
	Holder struct {
		dig.In
		PaymentService PaymentService
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewPaymentService); err != nil {
		return err
	}

	return nil
}
