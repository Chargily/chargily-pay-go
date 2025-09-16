package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// GetCustomer demonstrates how to retrieve an existing customer's details using the Chargily SDK.
func GetCustomer(customerID string, client *chargily.Client) {
	
	// Call the Get method on the Customers service to retrieve the customer by ID.
	customer, err := client.Customers.Get(customerID)
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error getting customer: %v\n", err)
		return
	}
	// Output the successfully retrieved customer details.
	fmt.Printf("Customer retrieved successfully: %+v\n", *customer)
}
