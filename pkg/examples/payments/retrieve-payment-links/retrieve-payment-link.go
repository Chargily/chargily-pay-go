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


	// payment link id
	paymentLinkId := "your-payment-link-id" // make sure to be a valid payment link id

    // retrieve the payment link
    paymentLink, err := client.GetPaymentLink(paymentLinkId)
    if err!= nil {
        fmt.Printf("Error retrieving payment link by ID: %v\n", err)
        return
    }

    fmt.Printf("retrieved payment link: %+v\n", paymentLink)
}