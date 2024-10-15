package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// DeleteCustomer demonstrates how to delete an existing customer using the Chargily SDK.
func DeleteCustomer(customerID string , client * chargily.Client) {
	// Call the Delete method on the Customers service to delete the customer by ID.
	if err := client.Customers.Delete(customerID); err != nil {
		// Handle deletion error.
		fmt.Printf("Error deleting customer: %v\n", err)
		return
	}
	// Output a success message for deletion.
	fmt.Println("Customer deleted successfully.")
}
