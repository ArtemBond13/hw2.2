package main

import (
	"fmt"
	"github.com/ArtemBond13/hw2.2/pkg/card"
)

func main() {
	cardSvc := card.NewService("Tinkof")
	fmt.Printf("%v\n", cardSvc)
}