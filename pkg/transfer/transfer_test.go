package transfer

import (
	"github.com/ArtemBond13/hw2.2/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc           *card.Service
		PercentTransfer   float64
		MinAmountTransfer int64
	}
	type args struct {
		from   string
		to     string
		amount int64
	}
	cardSvc := card.NewService("Tinkoff Bank")
	cardSvc.Add(&card.Card{Balance: 100_000_00, Number: "1001"}, &card.Card{Balance: 1000_00, Number: "1002"},
		&card.Card{Balance: 200_00, Number: "1103"}, &card.Card{Balance: 300_00, Number: "1104"},
		&card.Card{Balance: 2000_00, Number: "9000"}, &card.Card{Balance: 3000_00, Number: "9999"},
		&card.Card{Balance: 200_00, Number: "1345"}, &card.Card{Balance: 300_00, Number: "1346"},
		&card.Card{Balance: 2000_00, Number: "4500"}, &card.Card{Balance: 300_00, Number: "4600"},
		&card.Card{Balance: 200_00, Number: "2345"}, &card.Card{Balance: 300_00, Number: "5432"},
	)
	var tests = []struct {
		name      string
		fields    fields
		args      args
		wantTotal int64
		wantOk    bool
	}{
		// TODO: Add test cases
		{name: "MyBankCard->MyBankCard, enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"100", "1002", 1000_00}, wantTotal: 101000, wantOk: true},
		{name: "MyBankCard->MyBankCard, not enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"1103", "1104", 1000_00}, wantTotal: 101000, wantOk: false},
		{name: "MyBankCard->OtherBankCard, enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"9000", "9991", 1000_00}, wantTotal: 101000, wantOk: true},
		{name: "MyBankCard->OtherBankCard, not enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"1345", "1343", 1000_00}, wantTotal: 101000, wantOk: false},
		{name: "OtherBankCard->MyBankCard, enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"4400", "4600", 1000_00}, wantTotal: 101000, wantOk: true},
		{name: "OtherBankCard->OtherBankCard, not enough money ", fields: fields{cardSvc, 0.5, 10_00},
			args: args{"4400", "1209", 1000_00}, wantTotal: 101000, wantOk: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:           tt.fields.CardSvc,
				PercentTransfer:   tt.fields.PercentTransfer,
				MinAmountTransfer: tt.fields.MinAmountTransfer,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
