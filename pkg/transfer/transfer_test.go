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
	cardSVC.Add(&card.Card{Id: 1, Balance: 2340_00, Number: "5106 2134 1234 3452"},
		&card.Card{Id: 2, Balance: 10000_00, Number: "5106 2134 9876 0926"},
		&card.Card{Id: 3, Balance: 1000_00, Number: "5106 2145 9876 0881"},
		&card.Card{Id: 4, Balance: 45000_00, Number: "5106 2134 9876 9075"},
		&card.Card{Id: 5, Balance: 3458_00, Number: "5106 2156 2672 3825"},
		&card.Card{Id: 6, Balance: 1000_00, Number: "5106 2134 6723 7666"},
		&card.Card{Id: 7, Balance: 1000_00, Number: "5106 2134 4562 6790"},
		&card.Card{Id: 8, Balance: 100_00, Number: "5106 2134 9876 5426"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "MyCardBank->MyCardBank, money enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 1234 3452", "5106 2134 9876 0926", 100_00}, want: 110_00, wantErr: false},
		{name: "MyCardBank->MyCardBank, money not enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2145 9876 0881", "5106 2134 9876 9075", 1000_00}, want: 1010_00, wantErr: true},
		{name: "MyCardBank->OtherCardBank, money enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2156 2672 3825", "5121 0612 4534 1206", 100_00}, want: 110_00, wantErr: false},
		{name: "MyCardBank->OtherCardBank, money not enough", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 6723 7666", "4567 1234 3456 7654", 1000_00}, want: 1010_00, wantErr: false},
		{name: "OtherCardBank->MyCardBank", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5121 0612 4534 1206", "5106 2134 4562 6790", 100_00}, want: 110_00, wantErr: false},
		{name: "OtherCardBank->OtherCardBank", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"3452 3234 7432 3238", "1234 5678 2746 1314", 100_00}, want: 110_00, wantErr: false},
		{name: "Source card does not belong to my cards ->OtherCardBank", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 5436", "1234 6543 2746 3469", 100_00}, want: 110_00, wantErr: false},
		{name: "Source Card my bank -> Target card does not belong to my cards", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 9876 5436", "5106 2156 2323", 100_00}, want: 110_00, wantErr: true},
		{name: "Source and Target cards is valid", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 1234 3452", "5106 2134 9876 0926", 100_00}, want: 110_00, wantErr: false},
		{name: "Source card is not valid", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 1234 3458", "5106 2134 9876 0926", 100_00}, want: 110_00, wantErr: true},
		{name: "Target card is not valid", fields: fields{cardSVC, 0.5, 10_00},
			args: args{"5106 2134 1234 3452", "5106 2134 9876 0922", 100_00}, want: 110_00, wantErr: true},
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
