package chargily

import (
	"strings"
)

// Client is the structure that holds the API key , the endpoint for the Chargily API and the development mode.
type Client struct {
    apiKey   		string
    endpoint 		string
    rs              RequestSenderI // rs: stands for RequestSender and used to send custom http requests
	mode            Mode
}



// NewClient initializes and returns a new Client with the given API key and endpoint.
func NewClient(apiKey , mode string) (*Client, error) {

    //Set the API base URL based on the mode provided
    var api_baseUrl string

    //Verify the mode parameter
    // Verify the mode parameter
    if mode == "prod" {
        api_baseUrl = ProdAPIBaseUrl
    } else if mode == "test" {
        api_baseUrl = TestAPIBaseUrl
    } else {
        return nil, ErrInvalidMode
    }


    //new request sender 
    requestSender := NewRequestSender(apiKey)
    
    //return the client with it's configurations
    return &Client{
        apiKey:   	apiKey,
        endpoint: 	api_baseUrl,
        rs:         requestSender,
		mode:       Mode(mode), //test: for testing/development stage , prod: for production applications
    }, nil 
}



//retrieve the balance example 
func (c *Client) GetBalance() (*Balance, error) {

    var balance Balance
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "balance"}, ""), nil, &balance)

    if err != nil {
        return nil, err
    }
	// Return the parsed balance object
	return &balance, nil
}



//=========== CUSTOMERS AREA ==============//

// create a new customer
func (c *Client) CreateCustomer(customer *CreateCustomerParams) (*Customer, error){

    var customerResp Customer
    //create new customer request with the customer data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "customers"}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}


// update the customer
func (c *Client) UpdateCustomer(customerID string, customer *CreateCustomerParams) (*Customer, error){

    var customerResp Customer
    //update the customer data request with the new updated customer data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "customers/", customerID}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}

// retrieve a costumer
func (c *Client) GetCustomer(customerID string) (*Customer, error) {

    var customer Customer
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "customers/",customerID }, ""), nil, &customer)

    if err != nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customer, nil
}


// delete a specific customer 
func (c *Client) DeleteCustomer(customerID string) error {

    //send the request 
    err := c.rs.SendRequest("DELETE",  strings.Join([]string{c.endpoint, "customers/", customerID}, ""), nil, nil)

    if err != nil {
        return err
    }
    // Return nil if the request was successful
    return nil
}


// retrieve all customers ( an array of customers )
func (c *Client) GetAllCustomers() (*RetrieveAll[Customer], error) {

    var customers RetrieveAll[Customer]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "customers"}, ""), nil, &customers)

    if err != nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customers, nil
}




//========= PRODUCTS AREA ===========//
///////////////////////////////////////


//create a new product
func (c *Client) CreateProduct(product *CreateProductParams) (*Product, error){

    var productResp Product
    //create new product request with the product data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "products"}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


// Update the product with it's unique ID
func (c *Client) UpdateProduct(productId string,product *CreateProductParams) (*Product, error){
    var productResp Product

    //update existing product request with the new product data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "products/",productId}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


//retrieve a product using its unique ID
func (c *Client) GetProduct(productId string) (*Product, error) {

    var product Product
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "products/", productId }, ""), nil, &product)

    if err != nil {
        return nil, err
    }
    // Return the parsed product object
    return &product, nil
}


// retrieve all products 
func (c *Client) GetAllProducts() (*RetrieveAll[Product], error) {

    var products RetrieveAll[Product]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "products"}, ""), nil, &products)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &products, nil
}


// delete a specific product
func (c *Client) DeleteProduct(productId string) error {

    //send the request 
    err := c.rs.SendRequest("DELETE",  strings.Join([]string{c.endpoint, "products/", productId}, ""), nil, nil)

    if err!= nil {
        return err
    }
    // Return nil if the request was successful
    return nil
}


// Retrieve a products's prices using its ID 
func (c *Client) GetProductPrices(productId string) (*RetrieveAll[ProductPrice], error) {

    var prices RetrieveAll[ProductPrice]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "products/", productId, "/prices"}, ""), nil, &prices)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product prices object
    return &prices, nil
}


//============PRICES AREA ==============//
//Create Price of a product for a specific product
func (c *Client) CreatePrice(productPrice  * ProductPriceParams) (*ProductPrice, error) {
    var price ProductPrice
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "prices"}, ""), productPrice, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
} 


// update the product price data (not the price itself as mentioned in the docs of Chargily) for a specific product
func (c *Client) UpdatePrice(productId string, Data * UpdatePriceMetaDataParams ) (*ProductPrice, error) {
    var price ProductPrice
    //send the request 

    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "prices/", productId}, ""), Data , &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a price 
func (c *Client) GetPrice(productId string) (*ProductPrice, error) {

    var price ProductPrice
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "prices/", productId }, ""), nil, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a list of all prices available 
func (c *Client) GetAllPrices() (*RetrieveAll[ProductPrice], error) {

    var prices RetrieveAll[ProductPrice]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "prices"}, ""), nil, &prices)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &prices, nil
}



//============ CHECKOUT FUNCTIONALITIES =================// 

//create a checkout 
func (c *Client) CreateCheckout(checkout *CheckoutParams) (*Checkout, error) {
    var checkoutResp Checkout
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "checkouts"}, ""), checkout, &checkoutResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkoutResp, nil
}


// retrieve a checkout
func (c *Client) GetCheckout(checkoutId string) (*Checkout, error) {

    var checkout Checkout
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId }, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkout, nil
}


// retrieve all checkouts
func (c *Client) GetAllCheckouts() (*RetrieveAll[Checkout], error) {

    var checkouts RetrieveAll[Checkout]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts"}, ""), nil, &checkouts)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkouts, nil
}



// retrieve a checkout's items
func (c *Client) GetCheckoutItems(checkoutId string) (*RetrieveAll[CheckoutItems], error) {

    var items RetrieveAll[CheckoutItems]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId, "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout items object
    return &items, nil
}


// expires a checkout 
func (c *Client) ExpireCheckout(checkoutId string) (*Checkout ,error) {

    //send the request 
    var checkout Checkout
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId, "/expire"}, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return nil if the request was successful
    return &checkout ,nil
}



//================== PAYMENT LINKS =====================//

//create payment link
func (c *Client) CreatePaymentLink(paymentLink *CreatePaymentLinkParams) (*PaymentLink, error) {
    var link PaymentLink
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "payment-links"}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}

// update a Payment Link
func (c *Client) UpdatePaymentLink(paymentLinkId string, paymentLink *CreatePaymentLinkParams) (*PaymentLink, error) {
    var link PaymentLink
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "payment-links/", paymentLinkId}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve a payment link
func (c *Client) GetPaymentLink(paymentLinkId string) (*PaymentLink, error) {

    var link PaymentLink
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links/", paymentLinkId }, ""), nil, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve all payment links
func (c *Client) GetAllPaymentLinks() (*RetrieveAll[PaymentLink], error) {

    var links RetrieveAll[PaymentLink]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links"}, ""), nil, &links)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &links, nil
}



// retrieve a payment link's items
func (c *Client) GetPaymentLinkItems(productId string) (*RetrieveAll[PItemsData], error) {

    var items RetrieveAll[PItemsData]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links/", productId , "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link items object
    return &items, nil
}