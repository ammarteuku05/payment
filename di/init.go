package di

import (
	"payment/controller"
	"payment/internal/repository"
	"payment/internal/service"
	shared "payment/shared"
	"payment/shared/config"

	"go.uber.org/dig"
)

var (
	Container *dig.Container = dig.New()
)

func init() {
	if err := Container.Provide(config.New); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewOrm); err != nil {
		panic(err)
	}

	if err := shared.Register(Container); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}
}
