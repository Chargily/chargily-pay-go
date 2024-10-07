package retrievecustomers

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

	// Retrieve a single customer by ID
	customerID := "your-customer-id" // you should always provide a valid customer ID that already Exists
    customer, err := client.GetCustomer(customerID)
    if err!= nil {
        fmt.Printf("Error retrieving customer: %v\n", err)
        return
    }

    fmt.Printf("Retrieved customer: %+v\n", customer)
}