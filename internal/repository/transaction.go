package repository

import (
	"context"
	"payment/internal/entity"
	"payment/shared"

	"gorm.io/gorm"
)

type (
	// BeneficieryRepository is
	TransactionRepository interface {
		Insert(ctx context.Context, req entity.Transaction) error
		GetById(ctx context.Context, id string) (*entity.Transaction, bool, error)
		GetByIdStatusNull(ctx context.Context, id string) (*entity.Transaction, bool, error)
		UpdateStatus(ctx context.Context, db *gorm.DB, id string, status string) error
	}

	transactionImpl struct {
		shared.Deps
	}
)

// NewBeneficieryRepository is
func NewTransactionRepository(deps shared.Deps) TransactionRepository {
	return &transactionImpl{Deps: deps}
}

func (s *transactionImpl) Insert(ctx context.Context, req entity.Transaction) error {
	if err := s.DB.Create(&req).Error; err != nil {
		return err
	}

	return nil
}
func (s *transactionImpl) GetByIdStatusNull(ctx context.Context, id string) (*entity.Transaction, bool, error) {
	var (
		err         error
		transaction *entity.Transaction
	)

	if err = s.DB.Where("id = ? AND status is null", id).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}

	return transaction, true, nil
}

func (s *transactionImpl) GetById(ctx context.Context, id string) (*entity.Transaction, bool, error) {
	var (
		err         error
		transaction *entity.Transaction
	)

	if err = s.DB.Where("id = ?", id).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}

	return transaction, true, nil
}

func (repos *transactionImpl) UpdateStatus(ctx context.Context, db *gorm.DB, id string, status string) error {
	var (
		err         error
		transaction entity.Transaction
	)

	if err = db.Model(&transaction).Where("id = ? AND status IS NULL", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil

}
