package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// CreateProduct demonstrates how to create a new product using the Chargily SDK.
func CreateProduct(client *chargily.Client) {
	// Define the parameters for the new product.
	bodyRequestProduct := &models.CreateProductParams{
		Name:        "Test Product",
		Description: "This is a test product",
		Images:      []string{"valid-image-link"},
		Metadata:    map[string]any{"key": "value"},
	}

	// Call the Create method on the Products service to create the product.
	product, err := client.Products.Create(bodyRequestProduct)
	if err != nil {
		// Handle creation error.
		fmt.Printf("Error creating product: %v\n", err)
		return
	}
	// Output the details of the successfully created product.
	fmt.Printf("Product created successfully: %+v\n", product)
}
