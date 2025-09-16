package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// CreatePaymentLink demonstrates how to create a payment link using the Chargily SDK.
func CreatePaymentLink(client *chargily.Client, id string) {
	// Create items to be added to the payment link
	items := []models.PItems{
		{
			Price:              id,
			Quantity:           2,
			AdjustableQuantity: true,
		},
	}

	// Create the payment link parameters
	paymentLinkParams := &models.CreatePaymentLinkParams{
		Name:                   "Test Order for Payment Link",
		Items:                  items,
		AfterCompletionMessage: "message",
		Locale:                 "en",
		PassFeesToCustomer:     false,
		CollectShippingAddress: 1,
		Metadata: map[string]any{
			"order_id": "order_54321",
			"notes":    "This is a test order for payment link.",
		},
	}

	// Create the payment link using the client.
	paymentLink, err := client.PaymentLinks.Create(paymentLinkParams)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating payment link: %v\n", err)
		return
	}
	// Output the details of the created payment link.
	fmt.Printf("Payment link created successfully: %+v\n", paymentLink)
}
