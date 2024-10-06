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
            Price:              "price_2",
            Quantity:           2,
            AdjustableQuantity:  true,
        },
        {
            Price:              "price_7",
            Quantity:           1,
            AdjustableQuantity:  false,
        },
    }

    // Create the payment link parameters
    paymentLinkParams := &chargily.CreatePaymentLinkParams{
        Name:                   "Updated Order for Payment Link Testing",
        Items:                  items,
        AfterCompletionMessage: "Thank you for your order! see you soon",
        Locale:                 "en",
        PassFeesToCustomer:     true,
        CollectShippingAddress: true,
        Metadata: map[string]any{
            "order_id": "order_54321",
            "notes":    "This is an example order for payment link.",
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