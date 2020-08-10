package model

import (
	"reflect"
	"testing"
)

func TestNewAccount(t *testing.T) {
	type args struct {
		id      string
		balance float64
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{"", args{"170f5049-6a13-4a8e-afdb-fd168451b90e", -1}, NewAccount("170f5049-6a13-4a8e-afdb-fd168451b90e", -1)},
		{"", args{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, NewAccount("170f5049-6a13-4a8e-afdb-fd168451b90e", 0)},
		{"", args{"170f5049-6a13-4a8e-afdb-fd168451b90e", 100.1}, NewAccount("170f5049-6a13-4a8e-afdb-fd168451b90e", 100.1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAccount(tt.args.id, tt.args.balance); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_ID(t *testing.T) {
	type fields struct {
		id      string
		balance float64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"", 0}, ""},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, "170f5049-6a13-4a8e-afdb-fd168451b90e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				balance: tt.fields.balance,
			}
			if got := a.ID(); got != tt.want {
				t.Errorf("Account.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Balance(t *testing.T) {
	type fields struct {
		id      string
		balance float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", -1}, -1},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, 0},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 100.1}, 100.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				balance: tt.fields.balance,
			}
			if got := a.Balance(); got != tt.want {
				t.Errorf("Account.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Validate(t *testing.T) {
	type fields struct {
		id      string
		balance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty ID error", fields{"", 0}, true},
		{"negative balance error", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", -0.1}, true},

		// positive cases
		{"not empty ID and zero balance", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, false},
		{"not empty ID non zero balance", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 100.1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				balance: tt.fields.balance,
			}
			if err := a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Account.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_validateID(t *testing.T) {
	type fields struct {
		id      string
		balance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty ID error", fields{"", 0}, true},

		// positive cases
		{"not empty ID", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				balance: tt.fields.balance,
			}
			if err := a.validateID(); (err != nil) != tt.wantErr {
				t.Errorf("Account.validateID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_validateBalance(t *testing.T) {
	type fields struct {
		id      string
		balance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"negative balance error", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", -0.1}, true},

		// positive cases
		{"zero balance", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 0}, false},
		{"non zero balance", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", 100.1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				balance: tt.fields.balance,
			}
			if err := a.validateBalance(); (err != nil) != tt.wantErr {
				t.Errorf("Account.validateBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
