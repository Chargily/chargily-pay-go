package chargily

// Wallet represents the balance details for a specific currency.
type Wallet struct {
    Currency        string `json:"currency"`         // The currency of the wallet (e.g., "dzd", "usd", "eur")
    Balance         int64  `json:"balance"`          // The total balance in the wallet
    ReadyForPayout  string `json:"ready_for_payout"` // The amount ready for payout
    OnHold          int64  `json:"on_hold"`         // The amount on hold
}

// Balance represents the overall balance entity.
type Balance struct {
    Entity   string   `json:"entity"`   // The entity type (e.g., "balance")
    LiveMode bool     `json:"livemode"` // Indicates whether the mode is live
    Wallets  []Wallet `json:"wallets"`  // A slice of Wallet structs
}


// Address represents a customer's address.
type Address struct {
	Address     		string 						    `json:"address,omitempty"`
	City       			string 							`json:"city,omitempty"`
	State      			string 							`json:"state,omitempty"`
	ZipCode    			string 							`json:"zip_code,omitempty"`
	Country    			string 							`json:"country,omitempty"`
}



// CreateCustomerParams represents the parameters for creating a customer.
type CreateCustomerParams struct {
	Name     			string            				`json:"name,omitempty"`     // The name of the customer.
	Email    			string            				`json:"email,omitempty"`    // The email address of the customer.
	Phone    			string            				`json:"phone,omitempty"`    // The phone number of the customer.
	Address  			*Address          				`json:"address,omitempty"`  // The address of the customer.
	Metadata 			map[string]any    				`json:"metadata,omitempty"`  // Additional info about the customer.
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


type AllCustomersResponse struct {
	Livemode                bool                         `json:"livemode"`
	CurrentPage             int                          `json:"current_page"`
	Data                    []Customer                   `json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}



type CreateProductParams struct {
    Name                   string                        `json:"name,omitempty"`     // The name of the product.
    Description            string                        `json:"description,omitempty"`  // The description of the product.
    Images                 []string                      `json:"images,omitempty"`    // The URLs of images of the product, up to 8.
    Metadata               map[string]any                `json:"metadata,omitempty"`  // A set of key-value pairs for additional information about the product.
}


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


type AllProductsResponse struct {
	Livemode                bool                         `json:"livemode"`
	CurrentPage             int                          `json:"current_page"`
	Data                    []Product                    `json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}
 


// price operations related structs
type ProductPriceParams struct {
	Amount                  int64                        `json:"amount"`          // The price amount in cents.
	Currency                string                       `json:"currency"`         // The currency code (e.g., "dzd", "usd", "eur").
	ProductID               string                       `json:"product_id"`      // The ID of the product to which the price applies.
	Metadata                 map[string]any              `json:"metadata,omitempty"`  // Additional information about the price.
}


type ProductPrice struct {
	ID                      string                       `json:"id"`            // The unique identifier of the price.
	Entity                  string                       `json:"entity"`          // The entity type (e.g., "price").
	Livemode                bool                         `json:"livemode"`        // Indicates whether the mode is live.
	Amount                  int64                        `json:"amount"`          // The price amount in cents.
	Currency                string                       `json:"currency"`         // The currency code (e.g., "dzd", "usd", "eur").
	ProductID               string                       `json:"product_id"`      // The ID of the product to which the price applies.
	Metadata                map[string]any               `json:"metadata"`         // Additional information about the price.
	CreatedAt               int64                        `json:"created_at"`       // The timestamp of when the price was created.
	UpdatedAt               int64                        `json:"updated_at"`       // The timestamp of when the price was updated.
}


type AllPricesResponse struct {
	Livemode                bool                         `json:"livemode"`	      
	CurrentPage             int                          `json:"current_page"`
	Data                    []ProductPrice               `json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}




//Checkouts structs 


type Item struct {
	Price                  string                       `json:"price"`           // The ID of the price to be added to the checkout.
	Quantity               int                          `json:"quantity"`         // The quantity of the item.
}

// the checkout param to create a checkout
type CheckoutParams struct {
	Items                  []Item                        `json:"items"`           // The items to be added to the checkout.
	Amount                  int                          `json:"amount"`           // The total amount to be charged in cents.
	Currency                string                       `json:"currency"`         // The currency code (e.g., "dzd", "usd", "eur").
	PaymentMethod           string                       `json:"payment_method"`    // The payment method (e.g., "card", "cash").
	SuccessURL              string                       `json:"success_url"`      // The URL to redirect to after a successful checkout.
	CustomerID              string                       `json:"customer_id,omitempty"`  // The ID of the customer to be associated with the checkout.
	FailureURL              string                       `json:"failure_url,omitempty"`  // The URL to redirect to after a failed checkout.
	WebhookEndpoint         string                       `json:"webhook_endpoint,omitempty"`  // The URL to send a webhook to after the checkout.
	Description             string                       `json:"description,omitempty"`  // A description of the checkout.
	Locale                  string                       `json:"locale,omitempty"`  // The locale (e.g., "en", "fr", "es").
	ShippingAddress         string                       `json:"shipping_address,omitempty"`  // The shipping address to be associated with the checkout.
	CollectShippingAddress  bool                         `json:"collect_shipping_address,omitempty"`  // Indicates whether the shipping address should be collected.
	PercentageDiscount      int                          `json:"percentage_discount,omitempty"`  // The percentage discount to be applied to the checkout.
	AmountDiscount          int                          `json:"amount_discount,omitempty"`  // The amount discount to be applied to the checkout in cents.
	Metadata                map[string]any               `json:"metadata,omitempty"`  // Additional information about the checkout.
}



type Discount struct {
	Type  					string            `json:"type"`   // The type of discount (e.g., "percentage", "fixed").
	Value 					int               `json:"value"`  // The value of the discount.
}

// Checkout Object to be returned
type Checkout struct {
	ID                      string            `json:"id"`                          // The unique identifier of the checkout.
	Entity                  string            `json:"entity"`                      // The entity type (e.g., "checkout").
	Livemode                bool              `json:"livemode"`                    // Indicates whether the mode is live.
	Amount                  int64             `json:"amount"`                      // The total amount to be charged in cents.
	Currency                string            `json:"currency"`                    // The currency of the checkout (e.g., "dzd").
	Fees                    int64             `json:"fees"`                        // The fees charged to the checkout.
	FeesOnMerchant          int64             `json:"fees_on_merchant"`            // The fees charged to the merchant in cents.
	FeesOnCustomer          int64             `json:"fees_on_customer"`            // The fees charged to the customer in cents.
	PassFeesToCustomer      *bool             `json:"pass_fees_to_customer"`       // Indicates whether the fees should be passed to the customer. This can be null.
	ChargilyPayFeesAllocation string           `json:"chargily_pay_fees_allocation"` // The allocation of fees to Chargily Pay.
	Status                  string            `json:"status"`                      // The status of the checkout (e.g., "pending", "succeeded", "failed").
	Locale                  string            `json:"locale"`                      // The locale (e.g., "en", "fr", "es").
	Description             *string           `json:"description"`                 // A description of the checkout. This can be null.
	Metadata                *map[string]any   `json:"metadata"`                    // Additional information about the checkout. This can be null.
	SuccessURL              string            `json:"success_url"`                 // The URL to redirect to after a successful checkout.
	FailureURL              string            `json:"failure_url"`                 // The URL to redirect to after a failed checkout.
	WebhookEndpoint         *string           `json:"webhook_endpoint"`            // The URL to send a webhook to after the checkout. This can be null.
	PaymentMethod           *string           `json:"payment_method"`              // The payment method (e.g., "card", "cash"). This can be null.
	InvoiceID               *string           `json:"invoice_id"`                  // The ID of the invoice associated with the checkout. This can be null.
	CustomerID              string            `json:"customer_id"`                 // The ID of the customer associated with the checkout.
	PaymentLinkID           *string           `json:"payment_link_id"`             // The ID of the payment link associated with the checkout. This can be null.
	CreatedAt               int64             `json:"created_at"`                  // The timestamp of when the checkout was created.
	UpdatedAt               int64             `json:"updated_at"`                  // The timestamp of when the checkout was updated.
	ShippingAddress         *string           `json:"shipping_address"`            // The shipping address to be associated with the checkout. This can be null.
	CollectShippingAddress  bool              `json:"collect_shipping_address"`    // Indicates whether the shipping address should be collected.
	Discount               	Discount		  `json:"discount"` // The discount applied to the checkout.
	AmountWithoutDiscount   int64   		  `json:"amount_without_discount"`  // The amount without any discount.
	CheckoutURL             string  		  `json:"checkout_url"`             // The URL to access the checkout page.
}

 
type AllCheckoutsResponse struct {
	Livemode                bool                         `json:"livemode"`	      
	CurrentPage             int                          `json:"current_page"`
	Data                    []Checkout                   `json:"data"` 
	FirstPageURL            string                       `json:"first_page_url"`
	LastPage                int                          `json:"last_page"`
	LastPageURL             string                       `json:"last_page_url"`
	NextPageURL             *string                      `json:"next_page_url"` 
	Path                    string                       `json:"path"`
	PerPage                 int                          `json:"per_page"`
	PrevPageURL             *string                      `json:"prev_page_url"` 
	Total                   int                          `json:"total"`
}


type CheckoutItems struct {
	ID         				string            		      `json:"id"`                    // The unique identifier of the checkout item.
	Entity     				string            		      `json:"entity"`                // The entity type (e.g., "price").
	Amount     				int64             		      `json:"amount"`                // The amount of the checkout item in cents.
	Quantity   				int64             		      `json:"quantity"`              // The quantity of the checkout item.
	Currency   				string            		      `json:"currency"`              // The currency of the checkout item (e.g., "dzd").
	Metadata   				map[string]any    		      `json:"metadata,omitempty"`    // Optional metadata associated with the item.
	CreatedAt  				int64             		      `json:"created_at"`            // The timestamp when the item was created.
	UpdatedAt  				int64             		      `json:"updated_at"`            // The timestamp when the item was last updated.
	ProductID  				string            		      `json:"product_id"`            // The unique identifier of the associated product.
}



// all checkout's Items 
type AllCheckoutItems struct {
	Livemode     		    bool                          `json:"livemode"`          // Indicates whether the checkout is in live mode.
	CurrentPage  		    int                           `json:"current_page"`      // The current page of results.
	Data         		    []CheckoutItems                `json:"data"`              // A list of checkout items on the current page.
	FirstPageURL 		    string                        `json:"first_page_url"`    // The URL for the first page.
	LastPage     		    int                           `json:"last_page"`         // The last available page.
	LastPageURL  		    string                        `json:"last_page_url"`     // The URL for the last page.
	NextPageURL  		    *string                       `json:"next_page_url"`     // The URL for the next page, null if there is none.
	Path         		    string                        `json:"path"`              // The base path for the pagination.
	PerPage      		    int                           `json:"per_page"`          // The number of items per page.
	PrevPageURL  		    *string                       `json:"prev_page_url"`     // The URL for the previous page, null if there is none.
	Total        		    int                           `json:"total"`             // The total number of checkout items.
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
	CollectShippingAddress  int    						   `json:"collect_shipping_address"`   // Indicates whether the shipping address should be collected (0 or 1).
	URL                    string  						  `json:"url"`                         // The URL to access the payment link.
}


type Items struct {
	Price              		string 						   `json:"price"`                		// The price of the item as a string.
	Quantity           		int    						   `json:"quantity"`             		// The quantity of the item.
	AdjustableQuantity 		bool   						   `json:"adjustable_quantity"`  		// Indicates if the quantity is adjustable by the customer.
}

type CreatePaymentLinkParams struct {
	Name                   string            				`json:"name"`                         // The name associated with the order.
	Items                  []Item            				`json:"items"`                        // A list of items in the order.
	AfterCompletionMessage string            				`json:"after_completion_message"`     // A message displayed after order completion.
	Locale                 string            				`json:"locale"`                       // The locale (e.g., "en", "fr").
	PassFeesToCustomer     bool              				`json:"pass_fees_to_customer"`        // Indicates if fees are passed to the customer.
	CollectShippingAddress bool              				`json:"collect_shipping_address"`     // Indicates whether to collect a shipping address.
	Metadata               map[string]any    				`json:"metadata"`                     // Additional metadata for the order.
}
