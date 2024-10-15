package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

//=========== CUSTOMERS AREA ==============//
type Customers struct {
    client * Client
}

// create a new customer
func (c * Customers) Create(customer *models.CreateCustomerParams) (*models.Customer, error){

    var customerResp models.Customer
    //create new customer request with the customer data
    err := c.client.rs.SendRequest("POST",  strings.Join([]string{c.client.endpoint, "customers"}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}


// update the customer
func (c * Customers) Update(customerID string, customer *models.CreateCustomerParams) (*models.Customer, error){

    var customerResp models.Customer
    //update the customer data request with the new updated customer data
    err := c.client.rs.SendRequest("POST",  strings.Join([]string{c.client.endpoint, "customers/", customerID}, ""), customer, &customerResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customerResp, nil
}

// retrieve a costumer
func (c * Customers) Get(customerID string) (*models.Customer, error) {

    var customer models.Customer
    //send the request 
    err := c.client.rs.SendRequest("GET",  strings.Join([]string{c.client.endpoint, "customers/",customerID }, ""), nil, &customer)

    if err != nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customer, nil
}


// delete a specific customer 
func (c * Customers) Delete(customerID string) error {

    //send the request 
    err := c.client.rs.SendRequest("DELETE",  strings.Join([]string{c.client.endpoint, "customers/", customerID}, ""), nil, nil)

    if err != nil {
        return err
    }
    // Return nil if the request was successful
    return nil
}


// retrieve all customers ( an array of customers )
func (c * Customers) GetAll() (*models.RetrieveAll[models.Customer], error) {

    var customers models.RetrieveAll[models.Customer]
    //send the request 
    err := c.client.rs.SendRequest("GET",  strings.Join([]string{c.client.endpoint, "customers"}, ""), nil, &customers)

    if err != nil {
        return nil, err
    }
    // Return the parsed customer object
    return &customers, nil
}


