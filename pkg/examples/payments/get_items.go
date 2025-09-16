package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// RetrievePaymentLinkItems demonstrates how to retrieve items from a specific payment link using the Chargily SDK.
func RetrievePaymentLinkItems(client *chargily.Client, paymentLinkID string) {
	// Retrieve the items from the payment link using its ID.
	paymentLinkItems, err := client.PaymentLinks.GetItems(paymentLinkID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving payment link items: %v\n", err)
		return
	}
	// Output the details of the retrieved payment link items.
	fmt.Printf("Payment link items retrieved successfully: %+v\n", paymentLinkItems)
}
