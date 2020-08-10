package service

import (
	"github.com/MGerasimchuk/bet-api/internal/domain/model"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
)

type ChangeAccountBalanceService struct {
	transactionalChangeAccountBalanceRepoGenerator repository.TransactionalChangeAccountBalanceRepositoryGenerator
}

func NewChangeAccountBalanceService(transactionalChangeAccountBalanceRepoGenerator repository.TransactionalChangeAccountBalanceRepositoryGenerator) *ChangeAccountBalanceService {
	return &ChangeAccountBalanceService{transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator}
}

func (u *ChangeAccountBalanceService) IncreaseBalance(account *model.Account, transactionID string, transactionAmount float64, transactionSource string) error {
	transactionalChangeAccountBalanceUnitRepo := u.transactionalChangeAccountBalanceRepoGenerator()

	if err := transactionalChangeAccountBalanceUnitRepo.StartOperation(); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetAccountRepository().IncreaseBalance(account, transactionAmount); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	t := model.NewTransaction(transactionID, account.ID(), transactionAmount, model.TransactionStatusExecuted, transactionSource)
	if err := t.Validate(); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetTransactionRepository().Create(t); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	return transactionalChangeAccountBalanceUnitRepo.FinalOperation(nil)
}

func (u *ChangeAccountBalanceService) DecreaseBalance(account *model.Account, transactionID string, transactionAmount float64, transactionSource string) error {
	transactionalChangeAccountBalanceUnitRepo := u.transactionalChangeAccountBalanceRepoGenerator()

	if err := transactionalChangeAccountBalanceUnitRepo.StartOperation(); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetAccountRepository().DecreaseBalance(account, transactionAmount); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	t := model.NewTransaction(transactionID, account.ID(), -1*transactionAmount, model.TransactionStatusExecuted, transactionSource)
	if err := t.Validate(); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetTransactionRepository().Create(t); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	return transactionalChangeAccountBalanceUnitRepo.FinalOperation(nil)
}

func (u *ChangeAccountBalanceService) CancelTransaction(transaction *model.Transaction) error {
	transactionalChangeAccountBalanceUnitRepo := u.transactionalChangeAccountBalanceRepoGenerator()

	if err := transactionalChangeAccountBalanceUnitRepo.StartOperation(); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	account, err := transactionalChangeAccountBalanceUnitRepo.GetAccountRepository().FindByID(transaction.AccountID())
	if err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetAccountRepository().DecreaseBalance(account, transaction.Amount()); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	if err := transactionalChangeAccountBalanceUnitRepo.GetTransactionRepository().UpdateStatus(transaction, model.TransactionStatusCanceled); err != nil {
		return transactionalChangeAccountBalanceUnitRepo.FinalOperation(err)
	}

	return transactionalChangeAccountBalanceUnitRepo.FinalOperation(nil)
}
