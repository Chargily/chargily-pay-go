package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

//========= PRODUCTS AREA ===========//
///////////////////////////////////////


type Products struct {
    client * Client
}



//create a new product
func (p * Products) Create(product *models.CreateProductParams) (*models.Product, error){

    var productResp models.Product
    //create new product request with the product data
    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "products"}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


// Update the product with it's unique ID
func (p * Products) Update(productId string,product *models.CreateProductParams) (*models.Product, error){
    var productResp models.Product

    //update existing product request with the new product data
    err := p.client.rs.SendRequest("POST",  strings.Join([]string{p.client.endpoint, "products/",productId}, ""), product, &productResp)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &productResp, nil
}


//retrieve a product using its unique ID
func (p * Products) Get(productId string) (*models.Product, error) {

    var product models.Product
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "products/", productId }, ""), nil, &product)

    if err != nil {
        return nil, err
    }
    // Return the parsed product object
    return &product, nil
}


// retrieve all products 
func (p * Products) GetAll() (*models.RetrieveAll[models.Product], error) {

    var products models.RetrieveAll[models.Product]
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "products"}, ""), nil, &products)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product object
    return &products, nil
}


// delete a specific product
func (p * Products) Delete(productId string) error {

    //send the request 
    err := p.client.rs.SendRequest("DELETE",  strings.Join([]string{p.client.endpoint, "products/", productId}, ""), nil, nil)

    if err!= nil {
        return err
    }
    // Return nil if the request was successful
    return nil
}


// Retrieve a products's prices using its ID 
func (p * Products) GetPrices(productId string) (*models.RetrieveAll[models.ProductPrice], error) {

    var prices models.RetrieveAll[models.ProductPrice]
    //send the request 
    err := p.client.rs.SendRequest("GET",  strings.Join([]string{p.client.endpoint, "products/", productId, "/prices"}, ""), nil, &prices)

    if err!= nil {
        return nil, err
    }
    // Return the parsed product prices object
    return &prices, nil
}