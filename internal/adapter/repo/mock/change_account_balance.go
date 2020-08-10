package mock

import (
	"github.com/pkg/errors"

	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
)

type ChangeAccountBalanceRepository struct {
	actualAccounts     *[]*Account
	actualTransactions *[]*Transaction

	nonCommittedAccounts     *[]*Account
	nonCommittedTransactions *[]*Transaction

	accountRepo     repository.AccountRepository
	transactionRepo repository.TransactionRepository
}

func NewChangeAccountBalanceRepo(
	actualAccounts *[]*Account, actualTransactions *[]*Transaction,
	nonCommittedAccounts *[]*Account, nonCommittedTransactions *[]*Transaction,
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
) *ChangeAccountBalanceRepository {
	return &ChangeAccountBalanceRepository{
		actualAccounts:     actualAccounts,
		actualTransactions: actualTransactions,

		nonCommittedAccounts:     nonCommittedAccounts,
		nonCommittedTransactions: nonCommittedTransactions,

		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
	}
}

func (r *ChangeAccountBalanceRepository) StartOperation() error {
	return nil
}

func (r *ChangeAccountBalanceRepository) GetAccountRepository() repository.AccountRepository {
	return r.accountRepo
}

func (r *ChangeAccountBalanceRepository) GetTransactionRepository() repository.TransactionRepository {
	return r.transactionRepo
}

func (r *ChangeAccountBalanceRepository) FinalOperation(err error) error {
	if err == nil {
		*r.actualAccounts = *r.nonCommittedAccounts
		*r.actualTransactions = *r.nonCommittedTransactions
		return nil
	}

	*r.nonCommittedAccounts = *r.actualAccounts
	*r.nonCommittedTransactions = *r.actualTransactions

	return errors.Wrap(err, "Transaction rollbacked")
}
