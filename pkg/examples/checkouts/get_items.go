package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// RetrieveCheckoutItems demonstrates how to retrieve items from a specific checkout using the Chargily SDK.
func RetrieveCheckoutItems(client *chargily.Client, checkoutID string) {
	// Retrieve the items from the checkout using its ID.
	checkoutItems, err := client.Checkouts.GetItems(checkoutID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving checkout items: %v\n", err)
		return
	}
	// Output the details of the retrieved checkout items.
	fmt.Printf("Checkout items retrieved successfully: %+v\n", checkoutItems)
}
