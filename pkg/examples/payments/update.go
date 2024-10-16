package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// UpdatePaymentLink demonstrates how to update an existing payment link using the Chargily SDK.
func UpdatePaymentLink(client *chargily.Client, paymentLinkID string) {
	// Prepare updated payment link parameters.
	updatedPaymentLinkParams := &models.CreatePaymentLinkParams{
		Name:                   "Updated Test Order for Payment Link",
		Items:                  []models.PItems{
			{
				Price:              paymentLinkID,
				Quantity:           2,
				AdjustableQuantity: true,
			},
		},
		AfterCompletionMessage: "Thank you for your updated order!",
		Locale:                 "en",
		PassFeesToCustomer:     false,
		CollectShippingAddress: 0,
		Metadata: map[string]any{
			"order_id": "updated_order_54321",
			"notes":    "This is an updated test order for payment link.",
		},
	}

	// Update the payment link using the client.
	updatedPaymentLink, err := client.PaymentLinks.Update(paymentLinkID, updatedPaymentLinkParams)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating payment link: %v\n", err)
		return
	}
	// Output the details of the updated payment link.
	fmt.Printf("Payment link updated successfully: %+v\n", updatedPaymentLink)
}
