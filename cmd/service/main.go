package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.2/pkg/card"
)

func main() {
	cardSvc := card.NewService("Tinkof")

	fmt.Printf("%v\n", cardSvc)

	cardSvc.Add(&card.Card{Balance: 100_000_00, Number: "1001"})

	cardSvc.Add(&card.Card{Id: 1, Balance: 2340_00, Number: "2345"}, &card.Card{Id: 2, Balance: 10000_00, Number: "8945"})

	fmt.Println("Card is valid:", cardSvc.IsValidLunaAlgorithm("374652346956782346957823694857692364857368475368"))

	fmt.Println(cardSvc.FindByNumberMyService("5106 2145 1234 231"))
}
