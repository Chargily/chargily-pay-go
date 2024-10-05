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
