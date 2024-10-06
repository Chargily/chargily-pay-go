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
    items := []chargily.Items{
        {
            Price:              "price_1",
            Quantity:           2,
            AdjustableQuantity:  true,
        },
        {
            Price:              "price_2",
            Quantity:           1,
            AdjustableQuantity:  false,
        },
    }

    // Create the payment link parameters
    paymentLinkParams := &chargily.CreatePaymentLinkParams{
        Name:                   "Test Order for Payment Link",
        Items:                  items,
        AfterCompletionMessage: "Thank you for your order!",
        Locale:                 "en",
        PassFeesToCustomer:     false,
        CollectShippingAddress: true,
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