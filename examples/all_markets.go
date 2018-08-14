package main

import (
	"fmt"
	"log"
)

func main() {
	config := nomics.New("Your-API-Key")

	markets, err := nomics.Markets()
	if err != nil {
		log.Fatal(err)
	}

	for _, market := range markets {
		fmt.Println(market.Market, market.Quote)
	}
}
