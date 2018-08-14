package main

import (
	"fmt"
	"log"

	"github.com/kilpatty/nomics-go/nomics"
)

func main() {
	config := nomics.New("Your-API-Key")

	prices, err := nomics.Prices()
	if err != nil {
		log.Fatal(err)
	}

	for _, price := range prices {
		fmt.Println(price.Currency, price.Price)
	}
}
