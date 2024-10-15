package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// ListAllPaymentLinks demonstrates how to retrieve all payment links using the Chargily SDK.
func ListAllPaymentLinks(client *chargily.Client) {
	// Call the GetAll method on the PaymentLinks service to retrieve all payment links.
	paymentLinks, err := client.PaymentLinks.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all payment links: %v\n", err)
		return
	}
	// Output the list of successfully retrieved payment links.
	fmt.Printf("All payment links retrieved successfully: %+v\n", paymentLinks)
}
