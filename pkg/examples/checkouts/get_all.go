package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// ListAllCheckouts demonstrates how to retrieve all checkouts using the Chargily SDK.
func ListAllCheckouts(client *chargily.Client) {
	// Call the GetAll method on the Checkouts service to retrieve all checkouts.
	checkouts, err := client.Checkouts.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all checkouts: %v\n", err)
		return
	}
	// Output the list of successfully retrieved checkouts.
	fmt.Printf("All checkouts retrieved successfully: %+v\n", checkouts)
}
