package model

import (
	"errors"
)

type Account struct {
	id      string
	balance float64
}

func NewAccount(id string, balance float64) *Account {
	return &Account{id: id, balance: balance}
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Validate() error {
	if err := a.validateID(); err != nil {
		return err
	}
	if err := a.validateBalance(); err != nil {
		return err
	}

	return nil
}

func (a *Account) validateID() error {
	if a.id == "" {
		return errors.New("value of \"id\" field can't be empty")
	}

	return nil
}

func (a *Account) validateBalance() error {
	if a.balance < 0 {
		return errors.New("value of \"balance\" field should be greater than or equal zero")
	}

	return nil
}
