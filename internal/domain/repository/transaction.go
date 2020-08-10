package repository

import (
	"github.com/MGerasimchuk/bet-api/internal/domain/model"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	IsExists(id string) (bool, error)
	FindByID(id string) (*model.Transaction, error)
	GetLast(limit int) ([]*model.Transaction, error)
	UpdateStatus(transaction *model.Transaction, newStatus string) error
}
