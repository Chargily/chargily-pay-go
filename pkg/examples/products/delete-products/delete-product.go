package deleteproducts

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

	// the id of the product to update
	id := "your-product-id" // make sure to provide a valid id for the product

    // Update the product
    err = client.DeleteProduct(string(id))
    if err!= nil {
        fmt.Printf("Error updating product: %v\n", err)
        return
    }
    fmt.Print("Product Deleted successfully") 
}