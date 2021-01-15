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
	tinkof.Transfer(1, "0001", 10000_00)

}
