package createproducts

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

    // Create a new product
    product := &chargily.CreateProductParams{
        Name:        "New Product",
        Description: "This is a new product",
		Images: []string{"link1","link2","link3"},
		Metadata:  map[string]any{"key0": "value0","key1":"value1"},
    }

    newProduct, err := client.CreateProduct(product)
    if err!= nil {
        fmt.Printf("Error creating product: %v\n", err)
        return
    }
	fmt.Printf("Created product: %+v\n", newProduct)
}