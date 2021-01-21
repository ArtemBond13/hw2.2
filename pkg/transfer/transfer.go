package transfer

import (
	"errors"
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

var (
	ErrSourceCardInsufficientFunds = errors.New("sorry, there are not enough funds on the card")
	ErrSourceCardNotFound          = errors.New("this source card not found")
	ErrTargetCardNotFound          = errors.New("this target card not found")
	ErrSourceCardNotValid          = errors.New("this source card is not valid")
	ErrTargetCardNotValid          = errors.New("this target card is not valid")
)

// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int64) (int64, error) {
	total := int64(0)
	const (
		percentTransfer   = 0.5
		minAmountTransfer = 10_00
	)

	commission := int64(float64(amount/100) * percentTransfer)

	if commission < minAmountTransfer {
		commission = minAmountTransfer
	}
	total = amount + commission

	if !s.CardSvc.IsValidLunaAlgorithm(from) {
		return total, ErrSourceCardNotValid
	}

	if !s.CardSvc.IsValidLunaAlgorithm(to) {
		return total, ErrTargetCardNotValid
	}

	source, sourceErr := s.CardSvc.FindByNumber(from)
	target, targetErr := s.CardSvc.FindByNumber(to)

	if sourceErr != nil && targetErr != nil {
		return total, nil //fmt.Errorf("source error: %v, target error %v \n", source, target)
	}

	if sourceErr != nil {
		return total, nil //fmt.Errorf("source error: %v\n", source)
	}
	if targetErr != nil {
		return total, nil //fmt.Errorf("target error: %v\n", target)
	}

	if source != nil && source.Balance < total {
		return total, ErrSourceCardInsufficientFunds
	}
	if source != nil && target != nil {
		source.Balance -= total
		target.Balance += amount

	}
	return total, nil
}
