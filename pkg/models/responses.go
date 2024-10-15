package models

//=================G GENERIC RESPONSES =========================//
//////////////////////////////////////////////////////////////////
// Generic response for retrieving all entries of any type
type RetrieveAll[T any] struct {
	Livemode                bool                         `json:"livemode"`
	CurrentPage             int                          `json:"current_page"`
	Data                    []T                   		`json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}

/////////////////////////////////////////////////////////////////





///******************************************************************//
//====================== Responses and Types ========================//
// Balance represents the overall balance entity.
type Balance struct {
    Entity   				string   					`json:"entity"`   // The entity type (e.g., "balance")
    LiveMode 				bool     					`json:"livemode"` // Indicates whether the mode is live
    Wallets  				[]Wallet 					`json:"wallets"`  // A slice of Wallet structs
}


// Customer represents a customer entity.
type Customer struct {
    ID                      string                       `json:"id"`            // The unique identifier of the customer.
    Entity                  string                       `json:"entity"`          // The entity type (e.g., "customer")  
    Livemode                bool                         `json:"livemode"`        // Indicates whether the mode is live.
    Name                    string                       `json:"name"`             // The name of the customer.
    Email                   string                       `json:"email"`            // The email address of the customer.
    Phone                   string                       `json:"phone"`            // The phone number of the customer.
    Address                 *Address                     `json:"address"`          // The address of the customer.
    Metadata                map[string]any               `json:"metadata"`         // Additional info about the customer.
    UpdatedAt               int64                        `json:"updated_at"`       // The timestamp of when the customer was updates
    CreatedAt               int64                        `json:"created_at"`       // The timestamp of when the customer was created.
}


// Product Object 
type Product struct {
    ID                      string                       `json:"id"`            // The unique identifier of the product.
    Entity                  string                       `json:"entity"`          // The entity type (e.g., "product").
    Livemode                bool                         `json:"livemode"`        // Indicates whether the mode is live.
    Name                    string                       `json:"name"`             // The name of the product.
    Description             string                       `json:"description"`      // The description of the product.
    Images                  []string                     `json:"images"`          // The URLs of images of the product, up to 8.
    Metadata                map[string]any               `json:"metadata"`         // A set of key-value pairs for additional information about the product.
    CreatedAt               int64                        `json:"created_at"`       // The timestamp of when the product was created.
    UpdatedAt               int64                        `json:"updated_at"`       // The timestamp of when the product was updated.
}


// Checkout Object to be returned
type Checkout struct {
	ID                      string            			  `json:"id"`                          // The unique identifier of the checkout.
	Entity                  string            			  `json:"entity"`                      // The entity type (e.g., "checkout").
	Livemode                bool              			  `json:"livemode"`                    // Indicates whether the mode is live.
	Amount                  int64             			  `json:"amount"`                      // The total amount to be charged in cents.
	Currency                string            			  `json:"currency"`                    // The currency of the checkout (e.g., "dzd").
	Fees                    int64             			  `json:"fees"`                        // The fees charged to the checkout.
	FeesOnMerchant          int64             			  `json:"fees_on_merchant"`            // The fees charged to the merchant in cents.
	FeesOnCustomer          int64             			  `json:"fees_on_customer"`            // The fees charged to the customer in cents.
	PassFeesToCustomer      *bool             			  `json:"pass_fees_to_customer"`       // Indicates whether the fees should be passed to the customer. This can be null.
	ChargilyPayFeesAllocation string          			   `json:"chargily_pay_fees_allocation"` // The allocation of fees to Chargily Pay.
	Status                  string            			  `json:"status"`                      // The status of the checkout (e.g., "pending", "succeeded", "failed").
	Locale                  string            			  `json:"locale"`                      // The locale (e.g., "en", "fr", "es").
	Description             *string           			  `json:"description"`                 // A description of the checkout. This can be null.
	Metadata                *map[string]any   			  `json:"metadata"`                    // Additional information about the checkout. This can be null.
	SuccessURL              string            			  `json:"success_url"`                 // The URL to redirect to after a successful checkout.
	FailureURL              string            			  `json:"failure_url"`                 // The URL to redirect to after a failed checkout.
	WebhookEndpoint         *string           			  `json:"webhook_endpoint"`            // The URL to send a webhook to after the checkout. This can be null.
	PaymentMethod           *string           			  `json:"payment_method"`              // The payment method (e.g., "card", "cash"). This can be null.
	InvoiceID               *string           			  `json:"invoice_id"`                  // The ID of the invoice associated with the checkout. This can be null.
	CustomerID              string            			  `json:"customer_id"`                 // The ID of the customer associated with the checkout.
	PaymentLinkID           *string           			  `json:"payment_link_id"`             // The ID of the payment link associated with the checkout. This can be null.
	CreatedAt               int64             			  `json:"created_at"`                  // The timestamp of when the checkout was created.
	UpdatedAt               int64             			  `json:"updated_at"`                  // The timestamp of when the checkout was updated.
	ShippingAddress         *string           			  `json:"shipping_address"`            // The shipping address to be associated with the checkout. This can be null.
	CollectShippingAddress  int32             			  `json:"collect_shipping_address"`    // Indicates whether the shipping address should be collected.
	Discount               	Discount		  			  `json:"discount"` // The discount applied to the checkout.
	AmountWithoutDiscount   int64   		  			  `json:"amount_without_discount"`  // The amount without any discount.
	CheckoutURL             string  		  			  `json:"checkout_url"`             // The URL to access the checkout page.
}



// Payment Links 
type PaymentLink struct {
	ID                     string  						  `json:"id"`                          // The unique identifier of the payment link.
	Entity                 string  						  `json:"entity"`                      // The entity type (e.g., "payment_link").
	Livemode               bool    						  `json:"livemode"`                    // Indicates whether the mode is live.
	Name                   string  						  `json:"name"`                        // The name or description of the payment link.
	Active                 int     						  `json:"active"`                      // Indicates if the payment link is active (1) or inactive (0).
	AfterCompletionMessage string  						  `json:"after_completion_message"`    // A message displayed to the user after payment is completed.
	Locale                 string  						  `json:"locale"`                      // The locale (e.g., "ar", "en").
	PassFeesToCustomer      bool   						   `json:"pass_fees_to_customer"`      // Indicates whether the fees are passed to the customer.
	Metadata                map[string]any  						   `json:"metadata"`                   // Additional metadata associated with the payment link.
	CreatedAt              int64   						  `json:"created_at"`                  // The timestamp when the payment link was created.
	UpdatedAt              int64   						  `json:"updated_at"`                  // The timestamp when the payment link was last updated.
	CollectShippingAddress int32    						  `json:"collect_shipping_address"`   // Indicates whether the shipping address should be collected (0 or 1).
	URL                    string  						  `json:"url"`                         // The URL to access the payment link.
}
////////////////////////////////////////////////////////////////////
