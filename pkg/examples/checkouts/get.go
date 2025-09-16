package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// RetrieveCheckout demonstrates how to retrieve a specific checkout using the Chargily SDK.
func RetrieveCheckout(client *chargily.Client, checkoutID string) {
	// Retrieve the checkout using its ID.
	checkoutData, err := client.Checkouts.Get(checkoutID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving checkout: %v\n", err)
		return
	}
	// Output the details of the retrieved checkout.
	fmt.Printf("Checkout retrieved successfully: %+v\n", checkoutData)
}
