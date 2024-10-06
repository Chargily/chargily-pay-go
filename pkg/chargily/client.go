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
    //create new customer request with the customer data
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
func (c *Client) GetAllCustomers() (*AllCustomersResponse, error) {

    var customers AllCustomersResponse
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
    //create new product request with the product data
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
func (c *Client) GetAllProducts() (*AllProductsResponse, error) {

    var products AllProductsResponse
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
