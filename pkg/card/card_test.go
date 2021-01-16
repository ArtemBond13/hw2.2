package card

import "testing"

func TestService_IsValidMyCardBank(t *testing.T) {
	type fields struct {
		BankName string
		Cards    []*Card
	}
	type args struct {
		cardNumber string
	}
	cardSvc := NewService("Tinkof Bank")
	cardSvc.Add(&Card{Id: 1, Number: "5106 2134 7654 1234"})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Valid", fields: fields{"Tinkof Bank", []*Card{}}, args: args{
			"5106 2134 7654 1234"}, wantErr: false},
		{name: "not valid", fields: fields{"Tinkof Bank", []*Card{}}, args: args{
			"5106 2134 1234"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				BankName: tt.fields.BankName,
				Cards:    tt.fields.Cards,
			}
			if err := s.IsValidMyCardBank(tt.args.cardNumber); (err != nil) != tt.wantErr {
				t.Errorf("IsValidMyCardBank() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
