package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// UpdateProduct demonstrates how to update an existing product's details using the Chargily SDK.
func UpdateProduct(client *chargily.Client, productID string) {
	// Define the parameters for updating the product.
	bodyRequestProduct := &models.CreateProductParams{
		Name:        "Test Product of the Update",
		Description: "This is an updated test product",
		Images:      []string{"valid-image-link"},
		Metadata:    map[string]any{"key": "value"},
	}

	// Call the Update method on the Products service to update the product.
	product, err := client.Products.Update(productID, bodyRequestProduct)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating product: %v\n", err)
		return
	}
	// Output the details of the successfully updated product.
	fmt.Printf("Product updated successfully: %+v\n", product)
}
