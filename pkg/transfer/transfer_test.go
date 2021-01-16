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
	cardSVC := card.NewService("Tinkoff Bank")
	cardSVC.Add(&card.Card{Id: 1, Balance: 2340_00, Number: "2345"}, &card.Card{Id: 2, Balance: 10000_00, Number: "8945"},
		&card.Card{Id: 3, Balance: 1000_00, Number: "3096"}, &card.Card{Id: 4, Balance: 45000_00, Number: "9078"},
		&card.Card{Id: 5, Balance: 3458_00, Number: "8956"}, &card.Card{Id: 6, Balance: 1000_00, Number: "7645"},
		&card.Card{Id: 7, Balance: 1000_00, Number: "7812"}, &card.Card{Id: 8, Balance: 100_00, Number: "5436"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "MyCardBank->MyCardBank, money enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"2345", "8945", 100_00}, want: 110_00, wantErr: false},
		{name: "MyCardBank->MyCardBank, money not enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"3096", "9078", 1000_00}, want: 1010_00, wantErr: true},
		{name: "MyCardBank->OtherCardBank, money enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"8956", "1234", 100_00}, want: 110_00, wantErr: false},
		{name: "MyCardBank->OtherCardBank, money not enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"7645", "1234", 1000_00}, want: 1010_00, wantErr: false},
		{name: "OtherCardBank->MyCardBank", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"1234", "7812", 100_00}, want: 110_00, wantErr: true},
		{name: "OtherCardBank->OtherCardBank", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"9876", "1234", 100_00}, want: 1010_00, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:           tt.fields.CardSvc,
				PercentTransfer:   tt.fields.PercentTransfer,
				MinAmountTransfer: tt.fields.MinAmountTransfer,
			}
			got, err := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Card2Card() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Card2Card() got = %v, want %v", got, tt.want)
			}
		})
	}
}