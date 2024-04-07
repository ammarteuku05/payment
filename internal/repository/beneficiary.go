package repository

import (
	"context"
	"fmt"
	"payment/internal/entity"
	"payment/shared"

	"gorm.io/gorm"
)

type (
	// BeneficieryRepository is
	BeneficieryRepository interface {
		GetBeneficieryByAccountNumber(ctx context.Context, db *gorm.DB, accountNumber, beneficiaryBankId string) (*entity.Beneficiary, bool, error)
		UpdateAmount(ctx context.Context, db *gorm.DB, beneficiaryAccountNumber, beneficiaryBankId string, amount float64) error
	}

	beneficieryImpl struct {
		shared.Deps
	}
)

// NewBeneficieryRepository is
func NewBeneficieryRepository(deps shared.Deps) BeneficieryRepository {
	return &beneficieryImpl{Deps: deps}
}

// GetBeneficieryByAccountNumber is
func (repository *beneficieryImpl) GetBeneficieryByAccountNumber(ctx context.Context, db *gorm.DB, accountNumber, beneficiaryBankId string) (*entity.Beneficiary, bool, error) {
	var (
		beneficiery entity.Beneficiary
		err         error
	)

	if err = db.Joins("join beneficiary_banks bb on bb.id = beneficiaries.beneficiary_bank_id ").Preload("BeneficiaryBank").Where("beneficiaries.beneficiary_account_number = ?", accountNumber).First(&beneficiery).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &beneficiery, true, nil
}

func (repos *beneficieryImpl) UpdateAmount(ctx context.Context, db *gorm.DB, beneficiaryAccountNumber, beneficiaryBankId string, amount float64) error {
	var (
		err         error
		beneficiery entity.Beneficiary
	)

	if err = db.Model(&beneficiery).Where("beneficiary_account_number = ? AND beneficiary_bank_id = ?", beneficiaryAccountNumber, beneficiaryBankId).UpdateColumn("beneficiary_amount", amount).Error; err != nil {
		return err
	}

	fmt.Println(err)

	return nil

}
