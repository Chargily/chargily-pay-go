package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// RetrievePaymentLink demonstrates how to retrieve a specific payment link using the Chargily SDK.
func RetrievePaymentLink(client *chargily.Client, paymentLinkID string) {
	// Retrieve the payment link using its ID.
	paymentLinkData, err := client.PaymentLinks.Get(paymentLinkID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving payment link: %v\n", err)
		return
	}
	// Output the details of the retrieved payment link.
	fmt.Printf("Payment link retrieved successfully: %+v\n", paymentLinkData)
}
