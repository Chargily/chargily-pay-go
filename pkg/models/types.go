package models

//******************************************************************//
//=========================== COMMON TYPES =========================//
/////////////////////////////////////////////////////////////////////

// Address represents a customer's address.
type Address struct {
	Address     		string 							`json:"address,omitempty"`
	City       			string 							`json:"city,omitempty"`
	State      			string 							`json:"state,omitempty"`
	ZipCode    			string 							`json:"zip_code,omitempty"`
	Country    			string 							`json:"country,omitempty"`
}


// Wallet represents the balance details for a specific currency.
type Wallet struct {
    Currency        	string 							`json:"currency"`         // The currency of the wallet (e.g., "dzd", "usd", "eur")
    Balance         	int64  							`json:"balance"`          // The total balance in the wallet
    ReadyForPayout  	string 							`json:"ready_for_payout"` // The amount ready for payout
    OnHold          	int64  							`json:"on_hold"`         // The amount on hold
}


// ProductPrice represents a price entity.
type ProductPrice struct {
	ID                      string                       `json:"id"`            // The unique identifier of the price.
	Entity                  string                       `json:"entity"`          // The entity type (e.g., "price").
	Livemode                bool                         `json:"livemode"`        // Indicates whether the mode is live.
	Amount                  int64                        `json:"amount"`          // The price amount in cents.
	Currency                string                       `json:"currency"`         // The currency code (e.g., "dzd", "usd", "eur").
	ProductID               string                       `json:"product_id"`      // The ID of the product to which the price applies.
	Metadata                map[string]any               `json:"metadata,omitempty"`         // Additional information about the price.
	CreatedAt               int64                        `json:"created_at"`       // The timestamp of when the price was created.
	UpdatedAt               int64                        `json:"updated_at"`       // The timestamp of when the price was updated.
}


type CItems struct {
	Price                  string                       `json:"price"`           // The ID of the price to be added to the checkout.
	Quantity               int                          `json:"quantity"`         // The quantity of the item.
}


type Discount struct {
	Type  					string            `json:"type"`   // The type of discount (e.g., "percentage", "fixed").
	Value 					int               `json:"value"`  // The value of the discount.
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


type PItems struct {
	Price              		string 						   `json:"price"`                		// The price of the item as a string.
	Quantity           		int    						   `json:"quantity"`             		// The quantity of the item.
	AdjustableQuantity 		bool   						   `json:"adjustable_quantity"`  		// Indicates if the quantity is adjustable by the customer.
}


type PItemsData struct {
	ID                	string      						  `json:"id"`                  // The unique identifier of the item.
	Entity            	string      						  `json:"entity"`              // The entity type (e.g., "price").
	Amount            	int         						  `json:"amount"`              // The amount for the item.
	Quantity          	int         						  `json:"quantity"`            // The quantity of the item.
	AdjustableQuantity	int       						  	  `json:"adjustable_quantity"` // Indicates whether the quantity is adjustable (converted from 0 or 1 to bool).
	Currency          	string      						  `json:"currency"`            // The currency code (e.g., "dzd").
	Metadata          	interface{} 						  `json:"metadata"`            // Metadata associated with the item.
	CreatedAt         	int64       						  `json:"created_at"`          // Timestamp when the item was created.
	UpdatedAt         	int64       						  `json:"updated_at"`          // Timestamp when the item was last updated.
	ProductID         	string      						  `json:"product_id"`          // The associated product ID.
}

/////////////////////////////////////////////////////////////////////