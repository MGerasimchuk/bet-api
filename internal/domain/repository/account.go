package repository

import "github.com/MGerasimchuk/bet-api/internal/domain/model"

type AccountRepository interface {
	FindByID(id string) (*model.Account, error)
	IncreaseBalance(account *model.Account, amount float64) error
	DecreaseBalance(account *model.Account, amount float64) error
}
