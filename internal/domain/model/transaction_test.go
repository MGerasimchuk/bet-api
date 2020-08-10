package model

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTransaction(t *testing.T) {
	type args struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
	}
	tests := []struct {
		name string
		args args
		want *Transaction
	}{
		{"", args{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame}, NewTransaction("45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame)},
		{"", args{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer}, NewTransaction("45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTransaction(tt.args.id, tt.args.accountID, tt.args.amount, tt.args.status, tt.args.source)
			got.SetCreatedAt(tt.want.createdAt)
			got.SetUpdatedAt(tt.want.updatedAt)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_ID(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, "45224b58-a23a-4d51-997d-557ac78443dd"},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, "45224b58-a23a-4d51-997d-557ac78443dc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.ID(); got != tt.want {
				t.Errorf("Transaction.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_AccountID(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, "170f5049-6a13-4a8e-afdb-fd168451b90e"},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, "170f5049-6a13-4a8e-afdb-fd168451b90c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.AccountID(); got != tt.want {
				t.Errorf("Transaction.AccountID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Amount(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, 0},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.Amount(); got != tt.want {
				t.Errorf("Transaction.Amount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Status(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, TransactionStatusExecuted},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, TransactionStatusCanceled},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.Status(); got != tt.want {
				t.Errorf("Transaction.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_SetStatus(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	type args struct {
		status string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, args{TransactionStatusCanceled}},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, args{TransactionStatusExecuted}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			tr.SetStatus(tt.args.status)
			if got := tr.status; !reflect.DeepEqual(got, tt.args.status) {
				t.Errorf("Transaction.status = %v, want %v", got, tt.args.status)
			}
		})
	}
}

func TestTransaction_Source(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, TransactionSourceGame},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, TransactionSourceServer},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.Source(); got != tt.want {
				t.Errorf("Transaction.Source() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_CreatedAt(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, time.Unix(0, 0)},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, time.Unix(2, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.CreatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.CreatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_SetCreatedAt(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	type args struct {
		createdAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, args{time.Unix(4, 4)}},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, args{time.Unix(5, 5)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			tr.SetCreatedAt(tt.args.createdAt)
			if got := tr.createdAt; !reflect.DeepEqual(got, tt.args.createdAt) {
				t.Errorf("Transaction.createdAt = %v, want %v", got, tt.args.createdAt)
			}
		})
	}
}

func TestTransaction_UpdatedAt(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, time.Unix(1, 1)},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, time.Unix(3, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if got := tr.UpdatedAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transaction.UpdatedAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_SetUpdatedAt(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	type args struct {
		updatedAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dd", "170f5049-6a13-4a8e-afdb-fd168451b90e", 0, TransactionStatusExecuted, TransactionSourceGame, time.Unix(0, 0), time.Unix(1, 1)}, args{time.Unix(4, 4)}},
		{"", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, args{time.Unix(5, 5)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			tr.SetUpdatedAt(tt.args.updatedAt)
			if got := tr.updatedAt; !reflect.DeepEqual(got, tt.args.updatedAt) {
				t.Errorf("Transaction.updatedAt = %v, want %v", got, tt.args.updatedAt)
			}
		})
	}
}

func TestGetAvailableStatuses(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"", []string{TransactionStatusExecuted, TransactionStatusCanceled}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvailableStatuses(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableStatuses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAvailableSources(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"", []string{TransactionSourceGame, TransactionSourceServer, TransactionSourcePayment}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvailableSources(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableSources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransaction_Validate(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty ID", fields{"", "170f5049-6a13-4a8e-afdb-fd168451b90e", 100, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, true},
		{"empty accountID", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},
		{"zero amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 0, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},
		{"bad status", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, "bad-status", TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},
		{"bad source", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, "bad-source", time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_validateID(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty ID", fields{"", "170f5049-6a13-4a8e-afdb-fd168451b90e", 100, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.validateID(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.validateID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_validateAccountID(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty accountID", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "", 100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.validateAccountID(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.validateAccountID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_validateAmount(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"zero amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 0, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.validateAmount(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.validateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_validateStatus(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"bad status", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, "bad-status", TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.validateStatus(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.validateStatus() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransaction_validateSource(t *testing.T) {
	type fields struct {
		id        string
		accountID string
		amount    float64
		status    string
		source    string
		createdAt time.Time
		updatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"bad source", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100, TransactionStatusCanceled, "bad-source", time.Unix(2, 2), time.Unix(3, 3)}, true},

		// positive cases
		{"canceled server transaction with negative amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", -100, TransactionStatusCanceled, TransactionSourceServer, time.Unix(2, 2), time.Unix(3, 3)}, false},
		{"executed game transaction with positive amount", fields{"45224b58-a23a-4d51-997d-557ac78443dc", "170f5049-6a13-4a8e-afdb-fd168451b90c", 100.12, TransactionStatusExecuted, TransactionSourceGame, time.Unix(2, 2), time.Unix(3, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Transaction{
				id:        tt.fields.id,
				accountID: tt.fields.accountID,
				amount:    tt.fields.amount,
				status:    tt.fields.status,
				source:    tt.fields.source,
				createdAt: tt.fields.createdAt,
				updatedAt: tt.fields.updatedAt,
			}
			if err := tr.validateSource(); (err != nil) != tt.wantErr {
				t.Errorf("Transaction.validateSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
