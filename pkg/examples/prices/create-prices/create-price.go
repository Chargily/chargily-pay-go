package createprices

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
    if err!= nil {
        fmt.Printf("Error creating client: %v\n", err)
        return
    }

    // Create a new price
    price := &chargily.ProductPriceParams{
        ProductID: "your-product-id",
        Currency: "USD",
		Amount: 100,
		Metadata: map[string]any{"key":"value"},
    }

	// update the price metadata
    productPrice, err := client.CreatePrice(price)
    if err!= nil {
		fmt.Printf("Error creating price: %v\n", err)
	}


	fmt.Printf("Created price: %+v\n", productPrice)
}