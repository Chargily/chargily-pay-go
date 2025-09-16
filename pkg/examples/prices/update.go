package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)

// UpdatePrice demonstrates how to update an existing price's metadata using the Chargily SDK.
func UpdatePrice(client *chargily.Client, priceID string) {
	// Define the parameters for updating the price metadata.
	data := &models.UpdatePriceMetaDataParams{
		Metadata: map[string]any{"new_key": "new_value", "key3": "keykey"},
	}

	// Call the Update method on the Prices service to update the price.
	price, err := client.Prices.Update(priceID, data)
	if err != nil {
		// Handle update error.
		fmt.Printf("Error updating price: %v\n", err)
		return
	}
	// Output the details of the successfully updated price.
	fmt.Printf("Price updated successfully: %+v\n", price)
}
