package repository

import "go.uber.org/dig"

type (
	//Holder is
	Holder struct {
		dig.In

		BeneficieryRepository
		TransactionRepository
	}
)

// Register is
func Register(container *dig.Container) error {
	if err := container.Provide(NewBeneficieryRepository); err != nil {
		return err
	}

	if err := container.Provide(NewTransactionRepository); err != nil {
		return err
	}

	return nil
}
