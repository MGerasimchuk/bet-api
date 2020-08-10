package usecase

import (
	"errors"
	"fmt"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
	"github.com/MGerasimchuk/bet-api/internal/domain/service"
)

type TransactionUsecase struct {
	changeAccountBalanceService *service.ChangeAccountBalanceService
	transactionRepo             repository.TransactionRepository
}

func NewTransactionUsecase(
	changeAccountBalanceService *service.ChangeAccountBalanceService,
	transactionRepo repository.TransactionRepository,
) *TransactionUsecase {
	return &TransactionUsecase{
		changeAccountBalanceService: changeAccountBalanceService,
		transactionRepo:             transactionRepo,
	}
}

func (u *TransactionUsecase) CancelTransaction(transactionID string) error {
	transaction, err := u.transactionRepo.FindByID(transactionID)
	if err != nil {
		return err
	}
	if transaction == nil {
		return errors.New(fmt.Sprintf("transaction with id: \"%s\", is not found", transactionID))
	}

	if transaction.Status() == model.TransactionStatusCanceled {
		return errors.New(fmt.Sprintf("transaction with id: \"%s\", already cancelled", transaction.ID()))
	}

	err = u.changeAccountBalanceService.CancelTransaction(transaction)

	return err
}
