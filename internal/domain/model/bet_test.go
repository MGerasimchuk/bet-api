package model

import (
	"reflect"
	"testing"
)

func TestNewBet(t *testing.T) {
	type args struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name string
		args args
		want *Bet
	}{
		{"", args{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, NewBet("170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd")},
		{"", args{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, NewBet("170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBet(tt.args.accountID, tt.args.state, tt.args.amount, tt.args.transactionID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_AccountID(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, "170f5049-6a13-4a8e-afdb-fd168451b90e"},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, "170f5049-6a13-4a8e-afdb-fd168451b90c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.AccountID(); got != tt.want {
				t.Errorf("Bet.AccountID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_State(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, "bad-state"},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, BetStateWin},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.State(); got != tt.want {
				t.Errorf("Bet.State() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_Amount(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, 0},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.Amount(); got != tt.want {
				t.Errorf("Bet.Amount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_TransactionID(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, "45224b58-a23a-4d51-997d-557ac78443dd"},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, "45224b58-a23a-4d51-997d-557ac78443dc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.TransactionID(); got != tt.want {
				t.Errorf("Bet.TransactionID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_IsWin(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateLost, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, false},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.IsWin(); got != tt.want {
				t.Errorf("Bet.IsWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_IsLost(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateWin, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, false},
		{"", fields{"170f5049-6a13-4a8e-afdb-fd168451b90c", BetStateLost, 100, "45224b58-a23a-4d51-997d-557ac78443dc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if got := b.IsLost(); got != tt.want {
				t.Errorf("Bet.IsLost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAvailableStates(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{"", []string{BetStateWin, BetStateLost}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvailableStates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAvailableStates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBet_Validate(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty accountID error", fields{"", "win", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},
		{"incorrect state", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},
		{"zero amount", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},
		{"negative amount", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", -0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},
		{"empty transactionID", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", -0.1, ""}, true},

		// positive cases
		{"win bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"lost bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "lost", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if err := b.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Bet.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBet_validateAccountID(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty accountID error", fields{"", "win", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},

		// positive cases
		{"win bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"lost bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "lost", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if err := b.validateAccountID(); (err != nil) != tt.wantErr {
				t.Errorf("Bet.validateAccountID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBet_validateTransactionID(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"empty transactionID", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", -0.1, ""}, true},

		// positive cases
		{"win bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"lost bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "lost", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if err := b.validateTransactionID(); (err != nil) != tt.wantErr {
				t.Errorf("Bet.validateTransactionID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBet_validateState(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"incorrect state", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "bad-state", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},

		// positive cases
		{"win bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"lost bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "lost", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if err := b.validateState(); (err != nil) != tt.wantErr {
				t.Errorf("Bet.validateState() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBet_validateAmount(t *testing.T) {
	type fields struct {
		accountID     string
		state         string
		amount        float64
		transactionID string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// negative cases
		{"zero amount", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},
		{"negative amount", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", -0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, true},

		// positive cases
		{"win bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "win", 0.1, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
		{"lost bet", fields{"170f5049-6a13-4a8e-afdb-fd168451b90e", "lost", 100, "45224b58-a23a-4d51-997d-557ac78443dd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bet{
				accountID:     tt.fields.accountID,
				state:         tt.fields.state,
				amount:        tt.fields.amount,
				transactionID: tt.fields.transactionID,
			}
			if err := b.validateAmount(); (err != nil) != tt.wantErr {
				t.Errorf("Bet.validateAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
