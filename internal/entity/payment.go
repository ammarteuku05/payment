package entity

import (
	"sync"
	"time"
)

type BenefecieryStatus string

const (
	StatusActive   BenefecieryStatus = "active"
	StatusInActive BenefecieryStatus = "inactive"
)

type Beneficiary struct {
	ID                       string            `json:"id" gorm:"primaryKey"`
	BeneficiaryName          string            `json:"beneficiary_name"`
	BeneficiaryAccountNumber string            `json:"beneficiary_account_number"`
	BeneficiaryBank          BeneficiaryBank   `json:"beneficiary_bank"`
	BeneficiaryBankID        string            `json:"beneficiary_bank_id"`
	BeneficiaryAmount        float64           `json:"beneficiary_amount"`
	Status                   BenefecieryStatus `json:"status"`
	CreatedAt                time.Time         `json:"created_at"`
	UpdatedAt                time.Time         `json:"updated_at"`
	sync.Mutex
}

type BeneficiaryBank struct {
	ID                  string    `json:"id" gorm:"primaryKey"`
	BeneficiaryBankName string    `json:"beneficiary_bank_name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type Transaction struct {
	ID                     string    `json:"id" gorm:"primaryKey"`
	SenderAccountNumber    string    `json:"sender_account_number"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
	Amount                 float64   `json:"amount"`
	Status                 *string   `json:"status"`
	TransactionDate        string    `json:"transaction_date date"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}
