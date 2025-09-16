package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

//================== PAYMENT LINKS =====================//


type PaymentLinks struct {
    client * Client
}

//create payment link
func (p * PaymentLinks) Create(paymentLink *models.CreatePaymentLinkParams) (*models.PaymentLink, error) {
    var link models.PaymentLink
    //send the request 
    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "payment-links"}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}

// update a Payment Link
func (p * PaymentLinks) Update(paymentLinkId string, paymentLink *models.CreatePaymentLinkParams) (*models.PaymentLink, error) {
    var link models.PaymentLink
    //send the request 
    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "payment-links/", paymentLinkId}, ""), paymentLink, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve a payment link
func (p * PaymentLinks) Get(paymentLinkId string) (*models.PaymentLink, error) {

    var link models.PaymentLink
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "payment-links/", paymentLinkId }, ""), nil, &link)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &link, nil
}


// retrieve all payment links
func (p * PaymentLinks) GetAll() (*models.RetrieveAll[models.PaymentLink], error) {

    var links models.RetrieveAll[models.PaymentLink]
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "payment-links"}, ""), nil, &links)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link object
    return &links, nil
}



// retrieve a payment link's items
func (p * PaymentLinks) GetItems(productId string) (*models.RetrieveAll[models.PItemsData], error) {

    var items models.RetrieveAll[models.PItemsData]
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "payment-links/", productId , "/items"}, ""), nil, &items)

    if err!= nil {
        return nil, err
    }
    // Return the parsed payment link items object
    return &items, nil
}