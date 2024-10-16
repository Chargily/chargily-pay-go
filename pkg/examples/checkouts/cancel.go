package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// CancelCheckout demonstrates how to cancel a checkout using the Chargily SDK.
func CancelCheckout(client *chargily.Client, checkoutID string) {
	// Call the Expire method on the Checkouts service to cancel the checkout.
	checkoutEX, err := client.Checkouts.Expire(checkoutID)
	if err != nil {
		// Handle cancellation error.
		fmt.Printf("Error canceling checkout: %v\n", err)
		return
	}
	// Output the details of the successfully canceled checkout.
	fmt.Printf("Checkout canceled successfully: %+v\n", checkoutEX)
}
