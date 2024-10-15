package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// CreateCheckout demonstrates how to create a checkout using the Chargily SDK.
func CreateCheckout(client *chargily.Client , id string) {
	// Create example items to be added to the checkout.
	items := []models.CItems{
		{
			Price:    id,
			Quantity: 2,
		},
	}

	// Initialize the CheckoutParams struct with required fields.
	checkout := &models.CheckoutParams{
		Items:            items,
		PaymentMethod:    "edahabia",
		SuccessURL:       "https://your-site.com/success",
		FailureURL:       "https://your-site.com/failure",
		WebhookEndpoint:  "https://your-site.com/webhook",
		Description:      "Checkout for Order #12345",
		Locale:           "en",
		PercentageDiscount: 10,
	}

	// Create the checkout using the client.
	checkoutData, err := client.Checkouts.Create(checkout)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating checkout: %v\n", err)
		return
	}

	// Output the details of the created checkout.
	fmt.Printf("Checkout created successfully: %+v\n", checkoutData)
	fmt.Printf("Checkout Parameters: %+v\n", checkout)
}
