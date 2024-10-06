package updateproducts

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

    // new data 
    product := &chargily.CreateProductParams{
        Name:        "updated Product",
        Description: "This is an updates product",
		Images: []string{"link1","link2","link3","link4"},
		Metadata:  map[string]any{"key0": "value0"},
    }


	// the id of the product to update
	id := "your-product-id" // make sure to provide a valid id for the product

    // Update the product
    updatedProduct, err := client.UpdateProduct(string(id), product)
    if err!= nil {
        fmt.Printf("Error updating product: %v\n", err)
        return
    }
    fmt.Printf("Updated product: %+v\n", updatedProduct) 
}