package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

//============PRICES AREA ==============//

type Prices struct {
    client * Client
}

//Create Price of a product for a specific product
func (p * Prices) Create(productPrice  * models.ProductPriceParams) (*models.ProductPrice, error) {
    var price models.ProductPrice
    //send the request 
    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "prices"}, ""), productPrice, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
} 


// update the product price data (not the price itself as mentioned in the docs of Chargily) for a specific product
func (p * Prices) Update(productId string, Data * models.UpdatePriceMetaDataParams ) (*models.ProductPrice, error) {
    var price models.ProductPrice
    //send the request 

    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "prices/", productId}, ""), Data , &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a price 
func (p * Prices) Get(productId string) (*models.ProductPrice, error) {

    var price models.ProductPrice
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "prices/", productId }, ""), nil, &price)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &price, nil
}


// retrieve a list of all prices available 
func (p * Prices) GetAll() (*models.RetrieveAll[models.ProductPrice], error) {

    var prices models.RetrieveAll[models.ProductPrice]
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "prices"}, ""), nil, &prices)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product price object
    return &prices, nil
}
