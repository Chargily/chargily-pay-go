package reteriveallprices

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

	// create the product price and get back the price object
    prices, err := client.GetAllPrices()
    if err!= nil {
		fmt.Printf("Error while retrieving price: %v\n", err)
	}


	fmt.Printf("all prices: %+v\n", prices)
}