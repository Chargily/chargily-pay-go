package deletecustomers

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

    // Create a new customer
    customerID := "your-customer-id" // you should always provide a valid customer ID that already Exists
    err = client.DeleteCustomer(customerID)
    if err!= nil {
        fmt.Printf("Error deleting customer: %v\n", err)
        return
    }

    fmt.Printf("Deleted customer with ID: %s\n", customerID)
}