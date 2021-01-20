package card

import (
	"strconv"
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
func (s Service) FindByNumber(number string) (*Card, bool) {
	for _, card := range s.Cards {
		if card.Number == number {
			return card, true
		}
	}
	return nil, false
}

func (s Service) FindByNumberMyService(number string) (*Card, bool) {
	for _, card := range s.Cards {
		if strings.HasPrefix(card.Number, "5106 21") {
			if card.Number == number {
				return card, true
			}
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

// LunaAlgorithm возвращает ощибку если
func (s *Service) IsValidLunaAlgorithm(card string) bool {
	card = strings.ReplaceAll(card, " ", "")
	cardSlice := strings.Split(card, "")
	sNum := make([]int, len(cardSlice))
	sum1 := 0
	sum2 := 0
	for i, val := range cardSlice {
		sNum[i], _ = strconv.Atoi(val)
	}

	//fmt.Printf("Just array %d\n",sNum)
	for i, j := 0, len(sNum)-1; i < j; i, j = i+1, j-1 {
		sNum[i], sNum[j] = sNum[j], sNum[i]
	}
	//fmt.Printf("Reverse array %d\n",sNum)


	for i, num := range sNum {
		if i%2 != 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
			sum1 += num
		} else {
			sum2 += num
		}
	}

	totalSum := sum1 + sum2

	if totalSum%10 == 0 {
		//fmt.Println("Итоговая сумма", totalSum)
		return true
	} else {
		//fmt.Println("Итоговая сумма", totalSum)
		return false

	}
}
