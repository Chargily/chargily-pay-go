package retrieveallcustomers

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

    // Retrieve all customers
    customers, err := client.GetAllCustomers()
    if err!= nil {
        fmt.Printf("Error retrieving customers: %v\n", err)
        return
    }

    fmt.Printf("Retrieved customers: %+v\n", customers)
}