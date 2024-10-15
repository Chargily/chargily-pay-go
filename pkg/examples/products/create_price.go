package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// CreatePrice demonstrates how to create a price for a product using the Chargily SDK.
func CreatePrice(client *chargily.Client, productID string) {
	// Define the parameters for the new price.
	bodyRequestPrice := &models.ProductPriceParams{
		Amount:    10000,
		Currency:  "dzd",
		ProductID: productID, // Product ID to associate with this price
		Metadata:  map[string]any{"key": "value"},
	}

	// Call the Create method on the Prices service to create the price.
	price, err := client.Prices.Create(bodyRequestPrice)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating price: %v\n", err)
		return
	}
	// Output the details of the successfully created price.
	fmt.Printf("Price created successfully: %+v\n", price)
}
