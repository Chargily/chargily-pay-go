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
	itemss := []chargily.PItems{
		{
			Price:              "01j9k9m78jp18bdky07s0rxxtg",
			Quantity:           2,
			AdjustableQuantity:  true,
		},
	}

	// updated payment link data
	paymentLinkParams := &chargily.CreatePaymentLinkParams{
		Name:                   "Test Order for Payment Link",
		Items:                  itemss,
		AfterCompletionMessage: "Thank you for your order!",
		Locale:                 "en",
		PassFeesToCustomer:     true,
		CollectShippingAddress: 0,
		Metadata: map[string]any{
			"order_id": "order_54321",
			"notes":    "This is a test order for payment link.",
		},
	}


	// the id of payment id to be updated
	paymentLinkId:= "your-payment-link-id"

	// update 
	paymentLink, err := client.UpdatePaymentLink(paymentLinkId,paymentLinkParams)
    if err!= nil {
        fmt.Printf("Error updating payment link: %v\n", err)
        return
    }

    fmt.Printf("updated payment link: %+v\n", paymentLink)
}