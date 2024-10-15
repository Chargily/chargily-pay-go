package chargily

import (
	"github.com/Chargily/chargily-pay-go/pkg/chargily/utils"
)

// Client is the structure that holds the API key , the endpoint for the Chargily API and the development mode.
type Client struct {
    apiKey   		string
    endpoint 		string
    rs              utils.RequestSenderI // rs: stands for RequestSender and used to send custom http requests
	mode            Mode
    Balance         *Balance
    Customers       *Customers
    Prices          *Prices
    Products        *Products
    PaymentLinks    *PaymentLinks
    Checkouts       *Checkouts
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
        return nil, utils.ErrInvalidMode
    }


    //new request sender 
    requestSender := utils.NewRequestSender(apiKey)
    
    //return the client with it's configurations
    client :=  &Client{
        apiKey:   	apiKey,
        endpoint: 	api_baseUrl,
        rs:         requestSender,
		mode:       Mode(mode), //test: for testing/development stage , prod: for production applications
    }

    client.Balance =     &Balance{client: client}
    client.Customers =   &Customers{client: client}
    client.Prices =      &Prices{client: client}
    client.Products =    &Products{client: client}
    client.PaymentLinks = &PaymentLinks{client: client}
    client.Checkouts =   &Checkouts{client: client}

    return client, nil
}

