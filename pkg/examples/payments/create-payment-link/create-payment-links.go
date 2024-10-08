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


	// Create items to be added to the payment link
	itemss := []chargily.Items{
		{
			Price:              "01j9k9m78jp18bdky07s0rxxtg",
			Quantity:           2,
			AdjustableQuantity:  true,
		},
	}

	// Create the payment link parameters
	paymentLinkParams := &chargily.CreatePaymentLinkParams{
		Name:                   "Test Order for Payment Link",
		Items:                  itemss,
		AfterCompletionMessage: "Thank you for your order!",
		Locale:                 "en",
		PassFeesToCustomer:     false,
		CollectShippingAddress: 1,
		Metadata: map[string]any{
			"order_id": "order_54321",
			"notes":    "This is a test order for payment link.",
		},
	}

	// Create the payment link
	paymentLink, err := client.CreatePaymentLink(paymentLinkParams)
    if err!= nil {
        fmt.Printf("Error creating payment link: %v\n", err)
        return
    }

    fmt.Printf("created payment link: %+v\n", paymentLink)
}