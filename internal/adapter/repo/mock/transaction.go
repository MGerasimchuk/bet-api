package mock

import (
	"errors"
	"time"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
)

type TransactionRepository struct {
	Transactions *[]*Transaction
}

func (r *TransactionRepository) Create(transaction *model.Transaction) error {
	if t := r.findById(transaction.ID()); t != nil {
		return errors.New("transaction already exists")
	}

	*r.Transactions = append(*r.Transactions, fromDomainTransaction(transaction))

	return nil
}

func (r *TransactionRepository) IsExists(id string) (bool, error) {
	return r.findById(id) != nil, nil
}

func (r *TransactionRepository) FindByID(id string) (*model.Transaction, error) {
	t := r.findById(id)
	if t == nil {
		return nil, errors.New("transaction not found")
	}

	return toDomainTransaction(t), nil
}

func (r *TransactionRepository) GetLast(limit int) ([]*model.Transaction, error) {
	var from int
	if limit <= len(*r.Transactions) {
		from = len(*r.Transactions) - limit
	} else {
		from = 0
	}

	var domainTransactions []*model.Transaction
	for _, v := range (*r.Transactions)[from:len(*r.Transactions)] {
		domainTransactions = append(domainTransactions, toDomainTransaction(v))
	}

	return domainTransactions, nil
}

func (r *TransactionRepository) UpdateStatus(transaction *model.Transaction, newStatus string) error {
	t := r.findById(transaction.ID())
	if t == nil {
		return errors.New("transaction not found")
	}

	t.Status = newStatus

	return nil
}

func (r *TransactionRepository) findById(id string) *Transaction {
	for _, t := range *r.Transactions {
		if t.ID == id {
			return t
		}
	}

	return nil
}

func toDomainTransaction(t *Transaction) *model.Transaction {
	domainTransaction := model.NewTransaction(t.ID, t.AccountID, t.Amount, t.Status, t.Source)
	domainTransaction.SetCreatedAt(t.CreatedAt)
	domainTransaction.SetUpdatedAt(t.UpdatedAt)

	return domainTransaction
}

func fromDomainTransaction(domainTransaction *model.Transaction) *Transaction {
	return &Transaction{
		ID:        domainTransaction.ID(),
		AccountID: domainTransaction.AccountID(),
		Amount:    domainTransaction.Amount(),
		Status:    domainTransaction.Status(),
		Source:    domainTransaction.Source(),
		CreatedAt: domainTransaction.CreatedAt(),
		UpdatedAt: domainTransaction.UpdatedAt(),
	}
}

type Transaction struct {
	ID        string
	AccountID string
	Amount    float64
	Status    string
	Source    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
