package retrieveallpaymentlinks

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

	// retrieve all payment links
	paymentLinks, err := client.GetAllPaymentLinks()
    if err!= nil {
        fmt.Printf("Error retrieving all payment links: %v\n", err)
        return
    }

    fmt.Printf("payment-links: %+v\n", paymentLinks)
}