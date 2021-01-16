package card

import (
	"errors"
	"strings"
)

// Card абстракция банковской карты
type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

// Service хранятся карты банка
type Service struct {
	BankName string
	Cards    []*Card
}

// NewService возвращает сервис банка
func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

// Add добавляет карты в слайс Cards
func (s *Service) Add(cards ...*Card) {
	s.Cards = append(s.Cards, cards...)
}

// SearchByNumber поиска карты по номеру
func (s *Service) SearchByNumber(number string) *Card {
	for _, card := range s.Cards {
		if card.Number == number {
			return card
		}
	}
	return nil
}

//
func (s Service) FindByNumber(number string) (*Card, bool)  {
	for _, card := range s.Cards{
		if card.Number == number{
			return card, true
		}
	}
	return nil, false
}

func (s *Service) SearchById(id int64) *Card {
	for _, card := range s.Cards {
		if card.Id == id {
			return card
		}
	}
	return nil
}

func (s *Service) FindById(id int64) (*Card, bool) {
	for _, card := range s.Cards {
		if card.Id == id {
			return card, true
		}
	}
	return nil, false
}

// IssuerCard добавляет карту
func (s *Service) IssuerCard(id int64, issuer string, balance int64, number string) *Card {
	card := &Card{
		Id:       id,
		Issuer:   issuer,
		Balance:  balance,
		Currency: "RUB",
		Number:   number,
		Icon:     "http://...",
	}
	s.Cards = append(s.Cards, card)
	return card
}

var ErrMyCardNotValid = errors.New("there card not found my service")
// IsValidMyCardBank возвращает ошибку если карты нету в Сервисе 
func (s *Service) IsValidMyCardBank(cardNumber string) error  {
	prefix := "5106 21"
	for _, card := range s.Cards {
		if strings.HasPrefix(card.Number, prefix) == true{
			if card.Number != cardNumber {
				return  ErrMyCardNotValid
			}
		}
	}
	return nil
}