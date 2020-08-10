package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
)

type ChangeAccountBalanceRepository struct {
	transactionalDBInstance *gorm.DB
	accountRepo             repository.AccountRepository
	transactionRepo         repository.TransactionRepository
}

func NewChangeAccountBalanceRepo(
	transactionalDBInstance *gorm.DB,
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
) *ChangeAccountBalanceRepository {
	return &ChangeAccountBalanceRepository{
		transactionalDBInstance: transactionalDBInstance,
		accountRepo:             accountRepo,
		transactionRepo:         transactionRepo,
	}
}

// In this case(Postgres interface realisation), redundant
// But it can by useful in another interface realisation
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
		if commitError := r.transactionalDBInstance.Commit().Error; commitError != nil {
			if closeErr := r.transactionalDBInstance.Close(); closeErr != nil {
				return errors.Wrap(commitError, closeErr.Error())
			}

			return commitError
		}
		return nil
	}

	if rollbackErr := r.transactionalDBInstance.Rollback().Error; rollbackErr != nil {
		if closeErr := r.transactionalDBInstance.Close(); closeErr != nil {
			return errors.Wrap(rollbackErr, closeErr.Error())
		}

		return rollbackErr
	}

	return errors.Wrap(err, "transaction rollbacked")
}
