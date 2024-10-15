package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// GetAllPrices demonstrates how to retrieve all prices using the Chargily SDK.
func GetAllPrices(client *chargily.Client) {
	// Call the GetAll method on the Prices service to retrieve all prices.
	prices, err := client.Prices.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all prices: %v\n", err)
		return
	}
	// Output the list of successfully retrieved prices.
	fmt.Printf("All prices retrieved successfully: %+v\n", prices)
}
