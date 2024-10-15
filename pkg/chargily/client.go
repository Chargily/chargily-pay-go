package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
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
func (c *Client) GetBalance() (*models.Balance, error) {

    var balance models.Balance
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
func (c *Client) CreateCustomer(customer *models.CreateCustomerParams) (*models.Customer, error){

    var customerResp models.Customer
    //create new customer request with the customer data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "customers"}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}


// update the customer
func (c *Client) UpdateCustomer(customerID string, customer *models.CreateCustomerParams) (*models.Customer, error){

    var customerResp models.Customer
    //update the customer data request with the new updated customer data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "customers/", customerID}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}

// retrieve a costumer
func (c *Client) GetCustomer(customerID string) (*models.Customer, error) {

    var customer models.Customer
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
func (c *Client) GetAllCustomers() (*models.RetrieveAll[models.Customer], error) {

    var customers models.RetrieveAll[models.Customer]
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
func (c *Client) CreateProduct(product *models.CreateProductParams) (*models.Product, error){

    var productResp models.Product
    //create new product request with the product data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "products"}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


// Update the product with it's unique ID
func (c *Client) UpdateProduct(productId string,product *models.CreateProductParams) (*models.Product, error){
    var productResp models.Product

    //update existing product request with the new product data
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "products/",productId}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


//retrieve a product using its unique ID
func (c *Client) GetProduct(productId string) (*models.Product, error) {

    var product models.Product
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "products/", productId }, ""), nil, &product)

    if err != nil {
        return nil, err
    }
    // Return the parsed product object
    return &product, nil
}


// retrieve all products 
func (c *Client) GetAllProducts() (*models.RetrieveAll[models.Product], error) {

    var products models.RetrieveAll[models.Product]
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
func (c *Client) GetProductPrices(productId string) (*models.RetrieveAll[models.ProductPrice], error) {

    var prices models.RetrieveAll[models.ProductPrice]
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
func (c *Client) CreatePrice(productPrice  * models.ProductPriceParams) (*models.ProductPrice, error) {
    var price models.ProductPrice
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "prices"}, ""), productPrice, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
} 


// update the product price data (not the price itself as mentioned in the docs of Chargily) for a specific product
func (c *Client) UpdatePrice(productId string, Data * models.UpdatePriceMetaDataParams ) (*models.ProductPrice, error) {
    var price models.ProductPrice
    //send the request 

    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "prices/", productId}, ""), Data , &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a price 
func (c *Client) GetPrice(productId string) (*models.ProductPrice, error) {

    var price models.ProductPrice
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "prices/", productId }, ""), nil, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a list of all prices available 
func (c *Client) GetAllPrices() (*models.RetrieveAll[models.ProductPrice], error) {

    var prices models.RetrieveAll[models.ProductPrice]
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
func (c *Client) CreateCheckout(checkout *models.CheckoutParams) (*models.Checkout, error) {
    var checkoutResp models.Checkout
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "checkouts"}, ""), checkout, &checkoutResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkoutResp, nil
}


// retrieve a checkout
func (c *Client) GetCheckout(checkoutId string) (*models.Checkout, error) {

    var checkout models.Checkout
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId }, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkout, nil
}


// retrieve all checkouts
func (c *Client) GetAllCheckouts() (*models.RetrieveAll[models.Checkout], error) {

    var checkouts models.RetrieveAll[models.Checkout]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts"}, ""), nil, &checkouts)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkouts, nil
}



// retrieve a checkout's items
func (c *Client) GetCheckoutItems(checkoutId string) (*models.RetrieveAll[models.CheckoutItems], error) {

    var items models.RetrieveAll[models.CheckoutItems]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId, "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout items object
    return &items, nil
}


// expires a checkout 
func (c *Client) ExpireCheckout(checkoutId string) (*models.Checkout ,error) {

    //send the request 
    var checkout models.Checkout
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "checkouts/", checkoutId, "/expire"}, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return nil if the request was successful
    return &checkout ,nil
}



//================== PAYMENT LINKS =====================//

//create payment link
func (c *Client) CreatePaymentLink(paymentLink *models.CreatePaymentLinkParams) (*models.PaymentLink, error) {
    var link models.PaymentLink
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "payment-links"}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}

// update a Payment Link
func (c *Client) UpdatePaymentLink(paymentLinkId string, paymentLink *models.CreatePaymentLinkParams) (*models.PaymentLink, error) {
    var link models.PaymentLink
    //send the request 
    err := c.rs.SendRequest("POST",  strings.Join([]string{c.endpoint, "payment-links/", paymentLinkId}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve a payment link
func (c *Client) GetPaymentLink(paymentLinkId string) (*models.PaymentLink, error) {

    var link models.PaymentLink
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links/", paymentLinkId }, ""), nil, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve all payment links
func (c *Client) GetAllPaymentLinks() (*models.RetrieveAll[models.PaymentLink], error) {

    var links models.RetrieveAll[models.PaymentLink]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links"}, ""), nil, &links)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &links, nil
}



// retrieve a payment link's items
func (c *Client) GetPaymentLinkItems(productId string) (*models.RetrieveAll[models.PItemsData], error) {

    var items models.RetrieveAll[models.PItemsData]
    //send the request 
    err := c.rs.SendRequest("GET",  strings.Join([]string{c.endpoint, "payment-links/", productId , "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link items object
    return &items, nil
}