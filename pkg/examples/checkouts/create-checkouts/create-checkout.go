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


	//price id 
	priceID := "price-id" // must be an existing one 


	// Create example items to be added to the checkout.
	items := []chargily.CItems{
	    {
	        Price:    string(priceID),
	        Quantity: 2,
	    },
	}


	// Initialize the CheckoutParams struct with adjusted fields.
	checkoutParams := &chargily.CheckoutParams{
		Items: items,
		PaymentMethod:   "edahabia",
		SuccessURL:      "https://your-site.com/success",
		FailureURL:      "https://your-site.com/failure",
		WebhookEndpoint: "https://your-site.com/webhook",
		Description:     "Checkout for Order #12345",
		Locale:          "en",
		PercentageDiscount: 10,
	}

    // Create the checkout
    checkoutdata, err := client.CreateCheckout(checkoutParams)
    if err!= nil {
        fmt.Printf("Error creating checkout: %v\n", err)
        return
    }

	fmt.Printf("Created checkout: %+v\n", checkoutdata)
}