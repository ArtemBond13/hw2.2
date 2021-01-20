package card

import "testing"

func TestService_IsValidLunaAlgorithm(t *testing.T) {
	type fields struct {
		BankName string
		Cards    []*Card
	}
	type args struct {
		card string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "Card is Valid",
			args: args{"374652346956782346957823694857692364857368475368"},
			want: false,
		},
		{
			name: "Card is not Valid",
			args: args{"374652346956782346957823694857692364857387456834"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				BankName: tt.fields.BankName,
				Cards:    tt.fields.Cards,
			}
			if got := s.IsValidLunaAlgorithm(tt.args.card); got != tt.want {
				t.Errorf("IsValidLunaAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}
