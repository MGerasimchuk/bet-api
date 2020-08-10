package repository

type TransactionalChangeAccountBalanceRepository interface {
	StartOperation() error
	GetAccountRepository() AccountRepository
	GetTransactionRepository() TransactionRepository
	FinalOperation(err error) error
}
