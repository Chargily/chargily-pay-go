package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// CreateCustomer demonstrates how to create a new customer using the Chargily SDK.
func CreateCustomer(client * chargily.Client) {
	// Param example for a clear example using types.
	customerParams := &models.CreateCustomerParams{
		Name:  "John Doe", // Customer's name
		Email: "john.doe@example.com", // Customer's email
		Phone: "+1234567890", // Customer's phone number
		Address: &models.Address{
			Address: "123 Main St", // Street address
			State:   "NY",         // State
			Country: "US",         // Country
		},
	}

	// Call the Create method on the Customers service to create the customer.
	customer, err := client.Customers.Create(customerParams)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating customer: %v\n", err)
		return
	}
	// Output the successfully created customer details.
	fmt.Printf("Customer created successfully: %+v\n", *customer)
}
