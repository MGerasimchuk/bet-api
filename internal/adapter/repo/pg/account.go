package pg

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) FindByID(id string) (*model.Account, error) {
	a := &account{}
	err := r.db.Where(&account{ID: id}).First(a).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return toDomainAccount(a), nil
}

func (r *AccountRepository) IncreaseBalance(account *model.Account, amount float64) error {
	res := r.db.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", amount, account.ID())

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("increase balance error - account with id: \"%s\", is not found", account.ID()))
	}

	return nil
}

func (r *AccountRepository) DecreaseBalance(account *model.Account, amount float64) error {
	res := r.db.Exec("UPDATE accounts SET balance = balance - ? WHERE id = ? AND balance >= ?", amount, account.ID(), amount)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("decrease balance error - account with id: \"%s\", is not found, or balance of account less than decrese amount", account.ID()))
	}

	return nil
}

type account struct {
	ID      string  `gorm:"primary_key;column:id"`
	Balance float64 `gorm:"column:balance"`
}

func toDomainAccount(a *account) *model.Account {
	return model.NewAccount(a.ID, a.Balance)
}
