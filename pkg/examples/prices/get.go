package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// RetrievePrice demonstrates how to retrieve a specific price using the Chargily SDK.
func RetrievePrice(client *chargily.Client, priceID string) {
	// Retrieve the price by its ID.
	price, err := client.Prices.Get(priceID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving price: %v\n", err)
		return
	}
	// Output the details of the retrieved price.
	fmt.Printf("Price retrieved successfully: %+v\n", price)
}
