package pg

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/MGerasimchuk/bet-api/internal/domain/model"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) Create(domainTransaction *model.Transaction) error {
	t := fromDomainTransaction(domainTransaction)

	return r.db.Create(t).Error
}

func (r *TransactionRepository) IsExists(id string) (bool, error) {
	t := &transaction{}
	err := r.db.Where(&transaction{ID: id}).First(t).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *TransactionRepository) FindByID(id string) (*model.Transaction, error) {
	t := &transaction{}
	err := r.db.Where(&transaction{ID: id}).First(t).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return toDomainTransaction(t), nil
}

func (r *TransactionRepository) GetLast(limit int) ([]*model.Transaction, error) {
	var transactions []*transaction
	if err := r.db.Order("created_at desc").Limit(limit).Where(&transaction{Status: model.TransactionStatusExecuted}).Find(&transactions).Error; err != nil {
		return nil, err
	}

	var domainTransactions []*model.Transaction

	for i := range transactions {
		domainTransactions = append(domainTransactions, toDomainTransaction(transactions[i]))
	}

	return domainTransactions, nil
}

func (r *TransactionRepository) UpdateStatus(domainTransaction *model.Transaction, newStatus string) error {
	domainTransaction.SetStatus(newStatus)
	if err := domainTransaction.Validate(); err != nil {
		return err
	}

	t := fromDomainTransaction(domainTransaction)
	err := r.db.Model(&t).UpdateColumns(transaction{Status: t.Status, UpdatedAt: t.UpdatedAt}).Error

	return err
}

type transaction struct {
	ID        string    `gorm:"primary_key;column:id"`
	AccountID string    `gorm:"column:account_id"`
	Amount    float64   `gorm:"column:amount"`
	Status    string    `gorm:"column:status"`
	Source    string    `gorm:"column:source"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func toDomainTransaction(t *transaction) *model.Transaction {
	domainTransaction := model.NewTransaction(t.ID, t.AccountID, t.Amount, t.Status, t.Source)
	domainTransaction.SetCreatedAt(t.CreatedAt)
	domainTransaction.SetUpdatedAt(t.UpdatedAt)

	return domainTransaction
}

func fromDomainTransaction(t *model.Transaction) *transaction {
	return &transaction{ID: t.ID(),
		AccountID: t.AccountID(),
		Amount:    t.Amount(),
		Status:    t.Status(),
		Source:    t.Source(),
		CreatedAt: t.CreatedAt(),
		UpdatedAt: t.UpdatedAt(),
	}
}
