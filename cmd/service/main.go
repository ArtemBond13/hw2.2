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
	cardSvc.Add(&card.Card{Id: 1, Balance: 2340_00, Number: "5106 2145 1234 2312"}, &card.Card{Id: 2, Balance: 10000_00, Number: "8945"})
	tinkofTransfer := transfer.NewService(cardSvc, 0.5, 10)
	fmt.Println(tinkofTransfer.Card2Card("5106 2143 4523 1822",
		"5106 2145 1234 2312", 10_000_00))

	//fmt.Println(card.LunaAlgorithm("4561 2612 1234 5467"))
	fmt.Println(cardSvc.FindByNumberMyService("5106 2145 1234 23"))
}
