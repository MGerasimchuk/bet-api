package usecase

import (
	"errors"
	"fmt"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
	"github.com/MGerasimchuk/bet-api/internal/domain/repository"
	"github.com/MGerasimchuk/bet-api/internal/domain/service"
)

type BetUsecase struct {
	changeAccountBalanceService *service.ChangeAccountBalanceService
	accountRepo                 repository.AccountRepository
	transactionRepo             repository.TransactionRepository
}

func NewBetUsecase(
	changeAccountBalanceService *service.ChangeAccountBalanceService,
	accountRepo repository.AccountRepository,
	transactionRepo repository.TransactionRepository,
) *BetUsecase {
	return &BetUsecase{
		changeAccountBalanceService: changeAccountBalanceService,
		accountRepo:                 accountRepo,
		transactionRepo:             transactionRepo,
	}
}

// Actually I'm on the fence about the naming of this function and about "betResultSource" param name.
// I think, it's may be also transactionSource, and this usecase may probably has name like transaction use case, but because of
// state: win|lost field I think this usecase should be called "BetUsecase" and "service" in generally for me also about Bets
func (u *BetUsecase) BetResultProcess(betAccountID string, betAmount float64, betTransactionID string, betState string, betResultSource string) error {
	bet := model.NewBet(betAccountID, betState, betAmount, betTransactionID)
	if err := bet.Validate(); err != nil {
		return err
	}

	account, err := u.accountRepo.FindByID(bet.AccountID())
	if err != nil {
		return err
	}
	if account == nil {
		return errors.New(fmt.Sprintf("account with id: \"%s\", is not found", bet.AccountID()))
	}

	isExists, err := u.transactionRepo.IsExists(bet.TransactionID())
	if err != nil {
		return err
	}
	if isExists {
		return errors.New(fmt.Sprintf("transaction with id: \"%s\", already exists", bet.TransactionID()))
	}

	if bet.IsWin() {
		return u.changeAccountBalanceService.IncreaseBalance(account, bet.TransactionID(), bet.Amount(), betResultSource)
	}

	if bet.IsLost() {
		if account.Balance() < bet.Amount() {
			return errors.New(fmt.Sprintf("account with id: \"%s\", has balance, less than amount for decrease", bet.AccountID()))
		}
		return u.changeAccountBalanceService.DecreaseBalance(account, bet.TransactionID(), bet.Amount(), betResultSource)
	}

	return nil
}
