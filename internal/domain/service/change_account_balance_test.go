package service

import (
	"testing"
	"time"

	"github.com/MGerasimchuk/bet-api/internal/adapter/repo/mock"
	"github.com/MGerasimchuk/bet-api/internal/domain/model"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
)

func TestChangeAccountBalanceService_IncreaseBalance_AccountNotFoundError(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &actualAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &actualTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.IncreaseBalance(model.NewAccount("acc-not-found-id", 100), "tr2-id", 200, "game"); err == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, nil)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, nil)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccountBalance = not nil", a)
	}
	if a != nil && a.Balance() != 100 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 100)
	}

	//positive case
	//{
	//	"transaction already exists error",
	//		fields{transactionalChangeAccountBalanceRepoGenerator},
	//		args{model.NewAccount("acc1-id", 100), "tr1-id", 200, "game"}
	//},
	//
	//t.Run(tt.name, func(t *testing.T) {
	//	u := &ChangeAccountBalanceService{
	//		transactionalChangeAccountBalanceRepoGenerator: tt.fields.transactionalChangeAccountBalanceRepoGenerator,
	//	}
	//	if err := u.IncreaseBalance(tt.args.account, tt.args.transactionID, tt.args.transactionAmount, tt.args.transactionSource); (err != nil) != tt.wantErr {
	//		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, tt.wantErr)
	//	}
	//})
}

func TestChangeAccountBalanceService_IncreaseBalance_TransactionAlreadyExistsError(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &nonCommittedAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &nonCommittedTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.IncreaseBalance(model.NewAccount("acc1-id", 100), "tr2-id", 200, "incorrect-source"); err == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, true)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, false)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccount = not nil", a)
	}
	if a != nil && a.Balance() != 100 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 100)
	}
}

func TestChangeAccountBalanceService_IncreaseBalance_IncorrectSourceError(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &nonCommittedAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &nonCommittedTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.IncreaseBalance(model.NewAccount("acc1-id", 100), "tr2-id", 200, "bad-source"); err == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, true)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, nil)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccount = not nil", a)
	}
	if a != nil && a.Balance() != 100 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 100)
	}
}

func TestChangeAccountBalanceService_IncreaseBalance_Success(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &nonCommittedAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &nonCommittedTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.IncreaseBalance(model.NewAccount("acc1-id", 100), "tr2-id", 200, "game"); err != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, false)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, nil)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccount = not nil", a)
	}
	if a != nil && a.Balance() != 300 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 300)
	}
	transactionsCount, _ := transactionRepo.GetLast(10)
	if a != nil && len(transactionsCount) != 2 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() transactionCount = %v, wanttransactionCount = %v", len(transactionsCount), 300)
	}
}

func TestChangeAccountBalanceService_DecreaseBalance_BalanceLessThanAmount(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &nonCommittedAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &nonCommittedTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.DecreaseBalance(model.NewAccount("acc1-id", 100), "tr2-id", 200, "game"); err == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, true)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, nil)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccountBalance = not nil", a)
	}
	if a != nil && a.Balance() != 100 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 100)
	}
}

func TestChangeAccountBalanceService_DecreaseBalance_Success(t *testing.T) {
	actualAccounts := []*mock.Account{{"acc1-id", 100}}
	actualTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}
	nonCommittedAccounts := []*mock.Account{{"acc1-id", 100}}
	nonCommittedTransactions := []*mock.Transaction{{"tr1-id", "acc1-id", 100, "executed", "game", time.Unix(1, 1), time.Unix(2, 2)}}

	accountRepo := &mock.AccountRepository{Accounts: &nonCommittedAccounts}
	transactionRepo := &mock.TransactionRepository{Transactions: &nonCommittedTransactions}

	transactionalChangeAccountBalanceRepoGenerator := func() repository.TransactionalChangeAccountBalanceRepository {
		changeAccountBalanceRepo := mock.NewChangeAccountBalanceRepo(
			&actualAccounts, &actualTransactions,
			&nonCommittedAccounts, &nonCommittedTransactions,
			accountRepo, transactionRepo,
		)

		return changeAccountBalanceRepo
	}

	u := &ChangeAccountBalanceService{
		transactionalChangeAccountBalanceRepoGenerator: transactionalChangeAccountBalanceRepoGenerator,
	}
	if err := u.DecreaseBalance(model.NewAccount("acc1-id", 100), "tr2-id", 50, "game"); err != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", err, false)
	}

	a, e := accountRepo.FindByID("acc1-id")
	if e != nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() error = %v, wantErr %v", e, nil)
	}
	if a == nil {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() account = %v, wantAccountBalance = not nil", a)
	}
	if a != nil && a.Balance() != 50 {
		t.Errorf("ChangeAccountBalanceService.IncreaseBalance() accountBalance = %v, wantAccountBalance = %v", a.Balance(), 50)
	}
}
