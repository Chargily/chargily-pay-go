package expirescheckout

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

	// checkout id
	checkoutId := "your-checkout-id" // make sure to be a valid checkout id

    // expires the checkout
    checkout, err := client.ExpireCheckout(checkoutId)
    if err!= nil {
        fmt.Printf("Error expiring checkout: %v\n", err)
        return
    }

	fmt.Printf("expired checkout: %+v\n", checkout)
}