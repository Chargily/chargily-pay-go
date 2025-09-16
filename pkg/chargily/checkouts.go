package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

//============ CHECKOUT FUNCTIONALITIES =================//

type Checkouts struct {
    client * Client
}




//create a checkout
func (c * Checkouts) Create(checkout *models.CheckoutParams) (*models.Checkout, error) {
    var checkoutResp models.Checkout
    //send the request 
    err := c.client.rs.SendRequest("POST",  strings.Join([]string{c.client.endpoint, "checkouts"}, ""), checkout, &checkoutResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkoutResp, nil
}


// retrieve a checkout
func (c * Checkouts) Get(checkoutId string) (*models.Checkout, error) {

    var checkout models.Checkout
    //send the request 
    err := c.client.rs.SendRequest("GET",  strings.Join([]string{c.client.endpoint, "checkouts/", checkoutId }, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkout, nil
}


// retrieve all checkouts
func (c * Checkouts) GetAll() (*models.RetrieveAll[models.Checkout], error) {

    var checkouts models.RetrieveAll[models.Checkout]
    //send the request 
    err := c.client.rs.SendRequest("GET",  strings.Join([]string{c.client.endpoint, "checkouts"}, ""), nil, &checkouts)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout object
    return &checkouts, nil
}



// retrieve a checkout's items
func (c * Checkouts) GetItems(checkoutId string) (*models.RetrieveAll[models.CheckoutItems], error) {

    var items models.RetrieveAll[models.CheckoutItems]
    //send the request 
    err := c.client.rs.SendRequest("GET",  strings.Join([]string{c.client.endpoint, "checkouts/", checkoutId, "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed checkout items object
    return &items, nil
}


// expires a checkout 
func (c * Checkouts) Expire(checkoutId string) (*models.Checkout ,error) {

    //send the request 
    var checkout models.Checkout
    err := c.client.rs.SendRequest("POST",  strings.Join([]string{c.client.endpoint, "checkouts/", checkoutId, "/expire"}, ""), nil, &checkout)

    if err!= nil {
        return nil, err
    }
    // Return nil if the request was successful
    return &checkout ,nil
}
