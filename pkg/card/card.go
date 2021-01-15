package card

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
	Icon string
}

type Service struct {
	BankName string
	Cards []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) Add(cards ...*Card)  {
	s.Cards = append(s.Cards, cards...)
}