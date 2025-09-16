package chargily

import (
	"strings"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)


type Balance struct {
    client * Client
}

//retrieve the balance example
func (b * Balance) Get() (*models.Balance, error) {

    var balance models.Balance
    //send the request 
    err := b.client.rs.SendRequest("GET",  strings.Join([]string{b.client.endpoint, "balance"}, ""), nil, &balance)

    if err != nil {
        return nil, err
    }
	// Return the parsed balance object
	return &balance, nil
}