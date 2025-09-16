package models

//******************************************************************//
//=========================== PARAMS ===============================//
/////////////////////////////////////////////////////////////////////

// CreateCustomerParams represents the parameters for creating a customer.
type CreateCustomerParams struct {
	Name     			string            					`json:"name,omitempty"`     // The name of the customer.
	Email    			string            					`json:"email,omitempty"`    // The email address of the customer.
	Phone    			string            					`json:"phone,omitempty"`    // The phone number of the customer.
	Address  			*Address          					`json:"address,omitempty"`  // The address of the customer.
	Metadata 			map[string]any    					`json:"metadata,omitempty"`  // Additional info about the customer.
}


// price operations related structs
type ProductPriceParams struct {
	Amount                  int64                        	`json:"amount"`          // The price amount in cents.
	Currency                string                       	`json:"currency"`         // The currency code (e.g., "dzd", "usd", "eur").
	ProductID               string                       	`json:"product_id"`      // The ID of the product to which the price applies.
	Metadata                 map[string]any              	`json:"metadata,omitempty"`  // Additional information about the price.
}


// Creating a new product params 
type CreateProductParams struct {
    Name                   string                        	`json:"name,omitempty"`     // The name of the product.
    Description            string                        	`json:"description,omitempty"`  // The description of the product.
    Images                 []string                      	`json:"images,omitempty"`    // The URLs of images of the product, up to 8.
    Metadata               map[string]any                	`json:"metadata,omitempty"`  // A set of key-value pairs for additional information about the product.
}



//update PriceMetaData 
type UpdatePriceMetaDataParams struct {
    Metadata                 map[string]any              	`json:"metadata"`  // Additional information about the price.
}



// CheckoutParams represents the parameters required to create a checkout.
type CheckoutParams struct {
	Items                 	[]CItems          				`json:"items,omitempty"`               // Optional. The items to be added to the checkout.
	Amount                	int             				`json:"amount,omitempty"`                        // Required if items are not provided. The total amount in cents.
	Currency              	string          				`json:"currency,omitempty"`                      // Required if amount is provided. ISO currency code (e.g., "dzd").
	PaymentMethod         	string          				`json:"payment_method,omitempty"`      // Optional. Payment method (e.g., "edahabia", "cib").
	SuccessURL            	string          				`json:"success_url"`                   // Required. URL to redirect after successful payment.
	CustomerID            	string          				`json:"customer_id,omitempty"`         // Optional. ID of the existing customer.
	FailureURL            	string          				`json:"failure_url,omitempty"`         // Optional. URL to redirect after failed/canceled payment.
	WebhookEndpoint       	string          				`json:"webhook_endpoint,omitempty"`    // Optional. URL to receive webhook events after checkout.
	Description           	string          				`json:"description,omitempty"`         // Optional. Description of the checkout.
	Locale                	string          				`json:"locale,omitempty"`              // Optional. Checkout page language (e.g., "en", "fr", "ar").
	ShippingAddress       	string          				`json:"shipping_address,omitempty"`    // Optional. Customer's shipping address.
	CollectShippingAddress 	bool           					`json:"collect_shipping_address,omitempty"` // Optional. Whether to collect the shipping address from the customer.
	PercentageDiscount    	int             				`json:"percentage_discount,omitempty"` // Optional. Percentage discount, prohibited if amount discount is provided.
	AmountDiscount        	int             				`json:"amount_discount,omitempty"`     // Optional. Amount discount in cents, prohibited if percentage discount is provided.
	Metadata              	map[string]any  				`json:"metadata,omitempty"`            // Optional. Additional key-value pairs for extra checkout information.
}



// Create a new payment link 
type CreatePaymentLinkParams struct {
	Name                   string            				`json:"name"`                         // The name associated with the order.
	Items                  []PItems            				`json:"items"`                        // A list of items in the order.
	AfterCompletionMessage string            				`json:"after_completion_message"`     // A message displayed after order completion.
	Locale                 string            				`json:"locale"`                       // The locale (e.g., "en", "fr").
	PassFeesToCustomer     bool              				`json:"pass_fees_to_customer"`        // Indicates if fees are passed to the customer.
	CollectShippingAddress int32              				`json:"collect_shipping_address"`     // Indicates whether to collect a shipping address.
	Metadata               map[string]any    				`json:"metadata"`                     // Additional metadata for the order.
}

////////////////////////////////////////////////////////////////////
