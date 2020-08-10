package model

import (
	"errors"
	"strings"
	"time"
)

const (
	TransactionStatusExecuted = "executed"
	TransactionStatusCanceled = "canceled"

	TransactionSourceGame    = "game"
	TransactionSourceServer  = "server"
	TransactionSourcePayment = "payment"
)

type Transaction struct {
	id        string
	accountID string
	amount    float64
	status    string
	source    string
	createdAt time.Time
	updatedAt time.Time
}

func NewTransaction(id string, accountID string, amount float64, status string, source string) *Transaction {
	return &Transaction{
		id:        id,
		accountID: accountID,
		amount:    amount,
		status:    status,
		source:    source,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (t *Transaction) ID() string {
	return t.id
}

func (t *Transaction) AccountID() string {
	return t.accountID
}

func (t *Transaction) Amount() float64 {
	return t.amount
}

func (t *Transaction) Status() string {
	return t.status
}

func (t *Transaction) SetStatus(status string) {
	t.status = status
	t.updatedAt = time.Now()
}

func (t *Transaction) Source() string {
	return t.source
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Transaction) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}

func GetAvailableStatuses() []string {
	return []string{
		TransactionStatusExecuted,
		TransactionStatusCanceled,
	}
}

func GetAvailableSources() []string {
	return []string{
		TransactionSourceGame,
		TransactionSourceServer,
		TransactionSourcePayment,
	}
}

func (t *Transaction) Validate() error {
	if err := t.validateID(); err != nil {
		return err
	}
	if err := t.validateAccountID(); err != nil {
		return err
	}
	if err := t.validateAmount(); err != nil {
		return err
	}
	if err := t.validateStatus(); err != nil {
		return err
	}
	if err := t.validateSource(); err != nil {
		return err
	}

	return nil
}

func (t *Transaction) validateID() error {
	if t.id == "" {
		return errors.New("value of \"ID\" field can't be empty")
	}

	return nil
}

func (t *Transaction) validateAccountID() error {
	if t.accountID == "" {
		return errors.New("value of \"accountID\" field can't be empty")
	}

	return nil
}

func (t *Transaction) validateAmount() error {
	if t.amount == 0 {
		return errors.New("value of \"amount\" field can't be zero")
	}

	return nil
}

func (t *Transaction) validateStatus() error {
	availableStatuses := GetAvailableStatuses()
	for i := range availableStatuses {
		if t.status == availableStatuses[i] {
			return nil
		}
	}

	return errors.New("value of \"status\" field should be on of the following values: " + strings.Join(availableStatuses, ", "))
}

func (t *Transaction) validateSource() error {
	availableSources := GetAvailableSources()
	for i := range availableSources {
		if t.source == availableSources[i] {
			return nil
		}
	}

	return errors.New("value of \"source\" field should be on of the following values: " + strings.Join(availableSources, ", "))
}
