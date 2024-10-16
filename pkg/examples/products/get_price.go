package main

import (
	"fmt"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
)




func GetProductPrices(productId string, client * chargily.Client){
    // Get a product price
    prices, err := client.Products.GetPrices(productId)
    if err!= nil {
        fmt.Println("Error retrieving product price:", err)
    }

    fmt.Printf("Product prices : %v", prices)
}