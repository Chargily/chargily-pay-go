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