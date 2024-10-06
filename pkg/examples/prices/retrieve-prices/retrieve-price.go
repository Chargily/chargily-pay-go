package retrieveprices

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

	//product ID 
	prodId := "your-product-id"

	// create the product price and get back the price object
    productPrice, err := client.GetPrice(prodId)
    if err!= nil {
		fmt.Printf("Error while retrieving price: %v\n", err)
	}


	fmt.Printf("price: %+v\n", productPrice)
}