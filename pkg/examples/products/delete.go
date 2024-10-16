package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)

// DeleteProduct demonstrates how to delete an existing product using the Chargily SDK.
func DeleteProduct(client *chargily.Client, productID string) {
	// Call the Delete method on the Products service to delete the product by ID.
	if err := client.Products.Delete(productID); err != nil {
		// Handle deletion error.
		fmt.Printf("Error deleting product: %v\n", err)
		return
	}
	// Output a success message for deletion.
	fmt.Println("Product deleted successfully.")
}
