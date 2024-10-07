package retrievecheckouts

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


	//payment link ID 
	paymentLinkId  := "existing-payment-link-id" //must be unique and existing

	// retrieve a payment link's items
	paymentLinkItems, err := client.GetPaymentLinkItems(string(paymentLinkId))
    if err!= nil {
        fmt.Printf("Error getting payment link items: %v\n", err)
        return
    }
    fmt.Printf("Payment link items retrieved successfully: %+v\n\n\n\n", paymentLinkItems)
}