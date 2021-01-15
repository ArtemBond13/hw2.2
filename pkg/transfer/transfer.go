package transfer

import (
	"fmt"
	"github.com/ArtemBond13/hw2.2/pkg/card"
)

type Service struct {
	CardSvc           *card.Service
	PercentTransfer   float64
	MinAmountTransfer int64
}

func NewService(cardSVC *card.Service, percent float64, minAmount int64) *Service {
	return &Service{
		CardSvc:           cardSVC,
		PercentTransfer:   percent,
		MinAmountTransfer: minAmount,
	}
}

// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int64) (total int64, ok bool) {
	total = 0
	ofFrom := s.CardSvc.SearchByNumber(from)
	onTo := s.CardSvc.SearchByNumber(to)
	cardService := NewService(s.CardSvc, 0.5, 10_00)

	commission := int64(float64(amount/100) * cardService.PercentTransfer)
	if commission < cardService.MinAmountTransfer {
		commission = cardService.MinAmountTransfer
	}

	if ofFrom == nil && onTo == nil {
		total = amount + commission

		return total, true
	}

	if ofFrom == nil && onTo != nil {
		onTo.Balance += amount
		total = amount + commission
		fmt.Print(onTo.Balance, "\n")

		return total, true
	}

	if ofFrom != nil && onTo == nil {
		total = amount + commission
		if ofFrom.Balance < total {
			ok = false
			return total, ok
		}
		ofFrom.Balance -= total

		return total, true
	}
	total = amount + commission
	if ofFrom.Balance < total {
		return total, false
	}
	ofFrom.Balance -= total
	onTo.Balance += amount
	return total, true

}

func (s *Service) Transfer(fromId int64, toNumber string, amount int64) {
	source := s.CardSvc.SearchById(fromId)
	target := s.CardSvc.SearchByNumber(toNumber)

	source.Balance -= amount
	target.Balance += amount
}
