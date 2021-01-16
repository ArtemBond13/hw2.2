package transfer

import (
	"errors"
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

var ErrSourceCardInsufficientFunds = errors.New("sorry, there are not enough funds on the card")
// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int64) (int64, error) {
	total := int64(0)
	source := s.CardSvc.SearchByNumber(from)
	target := s.CardSvc.SearchByNumber(to)
	cardService := NewService(s.CardSvc, 0.5, 10_00)

	commission := int64(float64(amount/100) * cardService.PercentTransfer)
	if commission < cardService.MinAmountTransfer {
		commission = cardService.MinAmountTransfer
	}

	if source == nil && target == nil {
		total = amount + commission

		return total, nil
	}

	if source == nil && target != nil {
		target.Balance += amount
		total = amount + commission
		fmt.Print(target.Balance, "\n")

		return total, nil
	}

	if source != nil && target == nil {
		total = amount + commission
		if source.Balance < total {
			return total, ErrSourceCardInsufficientFunds
		}
		source.Balance -= total

		return total, nil
	}
	if source != nil && target !=nil {
		total = amount + commission
		if source.Balance < total {
			return total, ErrSourceCardInsufficientFunds
		}
		source.Balance -= total
		target.Balance += amount
		return total, nil
	}
	return total, nil
}

//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) {
//	if source, ok := s.CardSvc.FindById(fromId); ok{
//		if target, ok := s.CardSvc.FindByNumber(toNumber); ok {
//			source.Balance -= amount
//			target.Balance += amount
//
//		}
//
//	}
//}

// Make "early exit"
//func (s *Service) Transfer(fromId int64, toNumber string, amount int64) error{
//	source, ok := s.CardSvc.FindById(fromId)
//	if !ok {
//		err := TransferError("source card not found")
//		return err
//	}
//
//	target, ok := s.CardSvc.FindByNumber(toNumber)
//	if !ok {
//		err := TransferError("target card not found")
//		return err
//	}
//	source.Balance -= amount
//	target.Balance += amount
//	return nil
//}
//type TransferError string
//
//func (e TransferError) Error() string {
//	return string(e)

var (
	ErrSourceCardNotFound = errors.New("source card not found")
	ErrTargetCardNotFound = errors.New("target card not found")
)

func (s Service) Transfer(fromId int64, toNumber string, amount int64) error {
	source, ok := s.CardSvc.FindById(fromId)
	if !ok {
		return ErrSourceCardNotFound
	}

	target, ok := s.CardSvc.FindByNumber(toNumber)
	if !ok {
		return ErrTargetCardNotFound
	}

	source.Balance -= amount
	target.Balance += amount
	return nil
}
