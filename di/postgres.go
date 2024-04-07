package di

import (
	"fmt"
	"payment/shared/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	PostgresdbOption struct {
		ConnectionString                     string
		MaxLifeTimeConnection                time.Duration
		MaxIdleConnection, MaxOpenConnection int
	}
)

func NewPostgres(option *PostgresdbOption) (*gorm.DB, error) {
	var (
		opts = &gorm.Config{}
	)

	db, err := gorm.Open(postgres.Open(option.ConnectionString), opts)

	if err != nil {
		return nil, err
	}

	sql, err := db.DB()

	if err != nil {
		return nil, err
	}

	sql.SetConnMaxLifetime(option.MaxLifeTimeConnection)
	sql.SetMaxOpenConns(option.MaxOpenConnection)
	sql.SetMaxIdleConns(option.MaxIdleConnection)

	return db, nil
}

func NewOrm(config *config.Configuration) (*gorm.DB, error) {
	duration, err := time.ParseDuration(config.DbMaxLifeTimeConnection)

	if err != nil {
		return nil, err
	}

	DbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)
	var (
		opts = &PostgresdbOption{
			ConnectionString:      DbConnectionString,
			MaxOpenConnection:     config.DbMaxOpenConnection,
			MaxIdleConnection:     config.DbMaxIdleConnection,
			MaxLifeTimeConnection: duration,
		}
	)

	return NewPostgres(opts)
}
