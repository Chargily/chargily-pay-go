package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// GetAllCustomers demonstrates how to retrieve all customers using the Chargily SDK.
func GetAllCustomers(client * chargily.Client) {
	// Call the GetAll method on the Customers service to retrieve all customers.
	customers, err := client.Customers.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error getting all customers: %v\n", err)
		return
	}
	// Output the successfully retrieved list of customers.
	fmt.Printf("All customers retrieved successfully: %+v\n", customers)
}
