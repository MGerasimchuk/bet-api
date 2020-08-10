package mock

import (
	"errors"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
)

type AccountRepository struct {
	Accounts *[]*Account
}

func (r *AccountRepository) FindByID(id string) (*model.Account, error) {
	a := r.findById(id)
	if a == nil {
		return nil, errors.New("account not found")
	}

	return toDomainAccount(a), nil
}

func (r *AccountRepository) IncreaseBalance(account *model.Account, amount float64) error {
	a := r.findById(account.ID())
	if a == nil {
		return errors.New("account not found")
	}

	a.Balance += amount

	return nil
}

func (r *AccountRepository) DecreaseBalance(account *model.Account, amount float64) error {
	a := r.findById(account.ID())
	if a == nil {
		return errors.New("account not found")
	}
	if a.Balance < amount {
		return errors.New("balance is lower than amount for decrease")
	}

	a.Balance -= amount

	return nil
}

func (r *AccountRepository) findById(id string) *Account {
	for _, a := range *(r.Accounts) {
		if a.ID == id {
			return a
		}
	}

	return nil
}

func toDomainAccount(a *Account) *model.Account {
	return model.NewAccount(a.ID, a.Balance)
}

type Account struct {
	ID      string
	Balance float64
}
