package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// UpdateCustomer demonstrates how to update an existing customer's details using the Chargily SDK.
func UpdateCustomer(customerID string, client * chargily.Client) {

	// Define the parameters for updating the customer.
	customerParams := &models.CreateCustomerParams{
		Email: "john.updated.doe@example.com", // Updated email
		Address: &models.Address{
			Address: "1234 Main St", // Updated street address
			State:   "NYC",         // State
			Country: "USA",         // Country
		},
	}

	// Call the Update method on the Customers service to update the customer.
	customer, err := client.Customers.Update(customerID, customerParams)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating customer: %v\n", err)
		return
	}
	// Output the successfully updated customer details.
	fmt.Printf("Customer updated successfully: %+v\n", *customer)
}
