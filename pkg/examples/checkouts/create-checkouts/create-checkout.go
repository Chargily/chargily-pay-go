package createcheckouts

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


    // Create items to be added to the checkout
    items := []chargily.Item{
        {
            Price:    "price_1",
            Quantity: 2,
        },
        {
            Price:    "price_2",
            Quantity: 1,
        },
    }

    // Create the checkout parameters
    checkoutParams := &chargily.CheckoutParams{
        Items:                 items,
        Amount:                5000, 
        Currency:              "usd", 
        PaymentMethod:         "card", 
        SuccessURL:            "https://yourwebsite.com/success",
        CustomerID:            "customer_12345", // Optional
        FailureURL:            "https://yourwebsite.com/failure",
        WebhookEndpoint:       "https://yourwebsite.com/webhook",
        Description:           "Checkout for order #12345",
        Locale:                "en",
        ShippingAddress:       "123 Main St, Springfield, USA", // Optional
        CollectShippingAddress: true,
        PercentageDiscount:    10, // 10% discount
        AmountDiscount:        500, // $5.00 discount
        Metadata: map[string]any{
            "order_id": "order_12345",
            "notes":    "This is a test order",
        },
    }

    // Create the checkout
    checkout, err := client.CreateCheckout(checkoutParams)
    if err!= nil {
        fmt.Printf("Error creating checkout: %v\n", err)
        return
    }

	fmt.Printf("Created checkout: %+v\n", checkout)
}