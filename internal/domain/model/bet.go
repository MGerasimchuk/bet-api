package model

import (
	"errors"
	"strings"
)

const (
	BetStateWin  = "win"
	BetStateLost = "lost"
)

type Bet struct {
	accountID     string
	state         string
	amount        float64
	transactionID string
}

func NewBet(accountID string, state string, amount float64, transactionID string) *Bet {
	return &Bet{accountID: accountID, state: state, amount: amount, transactionID: transactionID}
}

func (b *Bet) AccountID() string {
	return b.accountID
}

func (b *Bet) State() string {
	return b.state
}

func (b *Bet) Amount() float64 {
	return b.amount
}

func (b *Bet) TransactionID() string {
	return b.transactionID
}

func (b *Bet) IsWin() bool {
	return b.state == BetStateWin
}

func (b *Bet) IsLost() bool {
	return b.state == BetStateLost
}

func GetAvailableStates() []string {
	return []string{
		BetStateWin,
		BetStateLost,
	}
}

func (b *Bet) Validate() error {
	if err := b.validateAccountID(); err != nil {
		return err
	}
	if err := b.validateTransactionID(); err != nil {
		return err
	}
	if err := b.validateState(); err != nil {
		return err
	}
	if err := b.validateAmount(); err != nil {
		return err
	}

	return nil
}

func (b *Bet) validateAccountID() error {
	if b.accountID == "" {
		return errors.New("value of \"accountID\" field can't be empty")
	}

	return nil
}

func (b *Bet) validateTransactionID() error {
	if b.transactionID == "" {
		return errors.New("value of \"TransactionID\" field can't be empty")
	}

	return nil
}

func (b *Bet) validateState() error {
	availableStates := GetAvailableStates()
	for i := range availableStates {
		if b.state == availableStates[i] {
			return nil
		}
	}

	return errors.New("value of \"state\" field should be on of the following values: " + strings.Join(availableStates, ", "))
}

func (b *Bet) validateAmount() error {
	if b.amount <= 0 {
		return errors.New("value of \"amount\" field should be greater than zero")
	}

	return nil
}
