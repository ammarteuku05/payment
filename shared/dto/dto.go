package dto

import (
	"payment/internal/entity"
	"time"
)

type BeneficieryResponse struct {
	BeneficiaryName          string                   `json:"beneficiary_name"`
	BeneficiaryAccountNumber string                   `json:"beneficiary_account_number"`
	BeneficiaryBank          entity.BeneficiaryBank   `json:"beneficiary_banks"`
	BeneficiaryAmount        float64                  `json:"beneficiary_amount"`
	Status                   entity.BenefecieryStatus `json:"status"`
	CreatedAt                time.Time                `json:"created_at"`
	UpdatedAt                time.Time                `json:"updated_at"`
}

type TransactionRequest struct {
	SenderAccountNumber        string  `json:"sender_account_number"`
	RecipientAccountNumber     string  `json:"recipient_account_number"`
	BeneficiaryBankIDSender    string  `json:"beneficiary_bank_id_sender"`
	BeneficiaryBankIDRecipient string  `json:"beneficiary_bank_id_recipient"`
	Amount                     float64 `json:"amount"`
}

type CallbackTransfer struct {
	SenderAccountNumber        string `json:"sender_account_number"`
	RecipientAccountNumber     string `json:"recipient_account_number"`
	BeneficiaryBankIDSender    string `json:"beneficiary_bank_id_sender"`
	BeneficiaryBankIDRecipient string `json:"beneficiary_bank_id_recipient"`
	TransferID                 string `json:"transfer_id"`
}
