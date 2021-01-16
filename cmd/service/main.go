package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.2/pkg/card"
	"github.com/ArtemBond13/hw2.2/pkg/transfer"
)

func main() {
	cardSvc := card.NewService("Tinkof")

	//c := cardSvc.SearchByNumber("0001")
	//fmt.Println(c.Balance)
	fmt.Printf("%v\n", cardSvc)

	cardSvc.Add(&card.Card{Balance: 100_000_00, Number: "1001"})
	tinkof := transfer.Service{CardSvc: cardSvc}
	err := tinkof.Transfer(1, "0001", 10000_00)
	if err != nil {
		switch err {
		case transfer.ErrSourceCardNotFound:
			fmt.Println("Sorry, can't complete transaction")
		case transfer.ErrTargetCardNotFound:
			fmt.Println("Please check target card number")
		default:
			fmt.Println("Something bad happened. try again later")
		}
	}
	cardSvc.Add(&card.Card{Id: 1, Balance: 2340_00, Number: "2345"}, &card.Card{Id: 2, Balance: 10000_00, Number: "8945"})
	tinkofTransfer := transfer.NewService(cardSvc, 0.5, 10)
	fmt.Println(tinkofTransfer.Card2Card("2345", "8945", 10_000_00))
}
