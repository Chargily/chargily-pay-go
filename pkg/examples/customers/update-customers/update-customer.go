package customers

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

    // Update a customer data
    customerParams := &chargily.CreateCustomerParams{
        Name:  "John Doe",
        Email: "john.doe@example.com",
        Phone: "+1 123 456 7890",
        Address: &chargily.Address{
            Address: "123 Main St",
            City:    "New York",
            State:   "NY",
            ZipCode: "10001",
            Country: "US",
        },
		Metadata: map[string]any{
            "custom_field": "value",
        },
	}

    // Get a customer by ID
    customerID := "your-customer-id" // you should always provide a valid customer ID that already Exists
    customer, err := client.UpdateCustomer(customerID,customerParams)
    if err!= nil {
        fmt.Printf("Error getting customer: %v\n", err)
        return
    }

    fmt.Printf("Updated customer: %+v\n", customer)
}