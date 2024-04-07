package shared

import (
	"payment/logger"

	"go.uber.org/dig"
	"gorm.io/gorm"

	"payment/shared/config"
)

type (
	Deps struct {
		dig.In

		DB     *gorm.DB
		Config *config.Configuration
		Logger logger.Logger

		CustomValidator *CustomValidator
	}
)

func Register(container *dig.Container) error {

	if err := container.Provide(NewCustomValidator); err != nil {
		return err
	}

	return nil
}
