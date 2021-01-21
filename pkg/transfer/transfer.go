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
)

var ErrSomethingBad = errors.New("something bad happened. Try again later")
// перевод денег с карты from на карту to в количестве amount
func (s *Service) Card2Card(from, to string, amount int64) (int64, error) {
	total := int64(0)
	cardService := NewService(s.CardSvc, 0.5, 10_00)

	commission := int64(float64(amount/100) * cardService.PercentTransfer)

	if commission < cardService.MinAmountTransfer {
		commission = cardService.MinAmountTransfer
	}
	total = amount + commission

	source, err := s.CardSvc.FindByNumberMyService(from)
	if err != nil {
		switch err {
		case card.ErrCardNotFoundMyService:
			return total, card.ErrCardNotFoundMyService
		case card.ErrCardNotOurBank:
			return total, nil
		default:
			return total, ErrSomethingBad
		}
	} else {
		if source.Balance < total {
			return total, ErrSourceCardInsufficientFunds
		}
		source.Balance -= total
	}

	target, err := s.CardSvc.FindByNumberMyService(to)
	if err != nil {
		switch err {
		case card.ErrCardNotFoundMyService:
			return total, card.ErrCardNotFoundMyService
		case card.ErrCardNotOurBank:
			return total, nil
		default:
			return total, ErrSomethingBad
		}
	} else {
		target.Balance += amount
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

func (s Service) Transfer(fromId int64, toNumber string, amount int64) error {
	source, ok := s.CardSvc.FindById(fromId)
	if !ok {
		return ErrSourceCardNotFound
	}

	target, err := s.CardSvc.FindByNumber(toNumber)
	if err != nil {
		return ErrTargetCardNotFound
	}

	source.Balance -= amount
	target.Balance += amount
	return nil
}
