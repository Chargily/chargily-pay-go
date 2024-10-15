package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// GetAllProducts demonstrates how to retrieve all products using the Chargily SDK.
func GetAllProducts(client *chargily.Client) {
	// Call the GetAll method on the Products service to retrieve all products.
	products, err := client.Products.GetAll()
	if err != nil {
		// Handle retrieval error.
		fmt.Printf("Error retrieving all products: %v\n", err)
		return
	}
	// Output the list of successfully retrieved products.
	fmt.Printf("All products retrieved successfully: %+v\n", products)
}
