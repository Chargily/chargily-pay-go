package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)



func main(){
	// Define your API key and mode
    apiKey := "your-api-key"
    mode := "test" // Could be "prod" or "test"

    // Create a new client instance
    client, err := chargily.NewClient(apiKey, mode)

	// error handling
	if err!= nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

	// Usage Example:
	client.Balance.Get()

    // Use the client to make API requests
    //...
    // Example usage:
	// client.Customers.(Get(), Update(...), Delete() ...).
	// client.Products.(Get(), Update(...), Delete() ...).
	// client.Prices
	// client.PaymentLinks.(Get(), Create(...), Delete() ...).
    // client.Checkouts.(Get(), Create(...), Delete() ...).
    //...
	// Read the documentation (README.md) for more information 
	// Or navigate to the examples folder
}