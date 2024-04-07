package service

import (
	"context"
	"fmt"
	"payment/internal/entity"
	"payment/internal/repository"
	"payment/shared"
	"payment/shared/dto"
	"payment/shared/errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type (
	// PaymentService is
	PaymentService interface {
		GetBeneficieryByAccountNumber(ctx context.Context, accountNumber, beneficiaryBankId string) (*dto.BeneficieryResponse, error)
		InsertTransfer(ctx context.Context, req *dto.TransactionRequest) (string, error)
		UpadateAmount(ctx context.Context, idTransfer, accountNumberRecipient, bankIdRecipient, accountNumberSender, bankIdSender string, wg *sync.WaitGroup) error
		CallBackUpdateStatusTransfer(ctx context.Context, req *dto.CallbackTransfer) error
	}

	paymentImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewPaymentService is
func NewPaymentService(repo repository.Holder, deps shared.Deps) PaymentService {
	return &paymentImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *paymentImpl) GetBeneficieryByAccountNumber(ctx context.Context, accountNumber, beneficiaryBankId string) (*dto.BeneficieryResponse, error) {
	res, check, err := s.repo.BeneficieryRepository.GetBeneficieryByAccountNumber(ctx, s.deps.DB, accountNumber, beneficiaryBankId)

	if !check {
		return nil, errors.ErrRecordNotFound
	}

	if err != nil {
		return nil, err
	}

	return &dto.BeneficieryResponse{
		BeneficiaryName:          res.BeneficiaryName,
		BeneficiaryBank:          res.BeneficiaryBank,
		BeneficiaryAccountNumber: res.BeneficiaryAccountNumber,
		BeneficiaryAmount:        res.BeneficiaryAmount,
		Status:                   res.Status,
		CreatedAt:                res.CreatedAt,
		UpdatedAt:                res.UpdatedAt,
	}, nil
}

func (s *paymentImpl) InsertTransfer(ctx context.Context, req *dto.TransactionRequest) (string, error) {
	DATE_MONTH_YEAR := "02 Jan 2006"
	_, checkRecipient, err := s.repo.BeneficieryRepository.GetBeneficieryByAccountNumber(ctx, s.deps.DB, req.RecipientAccountNumber, req.BeneficiaryBankIDRecipient)

	if !checkRecipient {
		return "", errors.ErrRecordNotFoundRecipient
	}

	if err != nil {
		return "", err
	}

	res, checkSender, err := s.repo.BeneficieryRepository.GetBeneficieryByAccountNumber(ctx, s.deps.DB, req.SenderAccountNumber, req.BeneficiaryBankIDRecipient)

	if !checkSender {
		return "", errors.ErrRecordNotFoundSender
	}

	if req.Amount > res.BeneficiaryAmount {
		return "", errors.ErrRecordTotal
	}

	if err != nil {
		return "", err
	}

	id := uuid.NewString()

	err = s.repo.Insert(ctx, entity.Transaction{
		ID:                     id,
		SenderAccountNumber:    req.SenderAccountNumber,
		RecipientAccountNumber: req.RecipientAccountNumber,
		Amount:                 req.Amount,
		TransactionDate:        time.Now().Format(DATE_MONTH_YEAR),
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
	})

	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *paymentImpl) UpadateAmount(ctx context.Context, idTransfer, accountNumberRecipient, bankIdRecipient, accountNumberSender, bankIdSender string, wg *sync.WaitGroup) error {
	// recipient
	resRecipient, check, err := s.repo.BeneficieryRepository.GetBeneficieryByAccountNumber(ctx, s.deps.DB, accountNumberRecipient, bankIdRecipient)

	if !check {
		return errors.ErrRecordNotFound
	}

	if err != nil {
		return err
	}

	defer wg.Done()
	resRecipient.Lock()
	defer resRecipient.Unlock()

	getById, checkTransfer, err := s.repo.TransactionRepository.GetById(ctx, idTransfer)

	if !checkTransfer {
		return errors.ErrRecordNotFoundTransfer
	}

	if err != nil {
		return err
	}

	amount := resRecipient.BeneficiaryAmount + getById.Amount

	err = s.repo.UpdateAmount(ctx, s.deps.DB, resRecipient.BeneficiaryAccountNumber, resRecipient.BeneficiaryBankID, amount)

	if err != nil {
		return err
	}

	// sender
	resSender, check, err := s.repo.BeneficieryRepository.GetBeneficieryByAccountNumber(ctx, s.deps.DB, accountNumberSender, bankIdSender)

	if !check {
		return errors.ErrRecordNotFound
	}

	if err != nil {
		return err
	}

	defer wg.Done()
	resSender.Lock()
	defer resSender.Unlock()

	getByIdSender, checkTransfer, err := s.repo.TransactionRepository.GetById(ctx, idTransfer)

	if !checkTransfer {
		return errors.ErrRecordNotFoundTransfer
	}

	if err != nil {
		return err
	}

	amount = resSender.BeneficiaryAmount - getByIdSender.Amount

	err = s.repo.UpdateAmount(ctx, s.deps.DB, resSender.BeneficiaryAccountNumber, resSender.BeneficiaryBankID, amount)

	if err != nil {
		return err
	}

	return nil
}

func (s *paymentImpl) CallBackUpdateStatusTransfer(ctx context.Context, req *dto.CallbackTransfer) error {
	response, check, err := s.repo.TransactionRepository.GetByIdStatusNull(ctx, req.TransferID)

	fmt.Printf("response %+v\n", response)
	fmt.Println(check)
	if !check {
		return errors.ErrRecordNotFoundTransfer
	}

	if err != nil {
		return err
	}

	status := "successed"
	err = s.repo.TransactionRepository.UpdateStatus(ctx, s.deps.DB, req.TransferID, status)

	if err != nil {
		return err
	}

	go func() {
		var wg sync.WaitGroup
		wg.Add(2)

		fmt.Println("cek disini")

		s.UpadateAmount(ctx, req.TransferID, req.RecipientAccountNumber, req.BeneficiaryBankIDRecipient, req.SenderAccountNumber, req.BeneficiaryBankIDSender, &wg)
	}()

	return nil
}
