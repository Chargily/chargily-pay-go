package retrieveallproducts

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

    // get all products 
    products, err := client.GetAllProducts()
    if err!= nil {
        fmt.Printf("Error updating product: %v\n", err)
        return
    }
    fmt.Printf("products : %+v\n", products.Data) 
}