package chargilypaygo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Chargily struct {
	baseUrl string
	apiKey  string
}

// Chargily client constructor
// Returns an error in case (mode) is neither 'live' nor 'test'
func NewChargily(mode, apiKey string) (*Chargily, error) {
	var apiMode string
	if mode == "test" {
		apiMode = "https://pay.chargily.net/test/api/v2"
	} else if mode == "live" {
		apiMode = "https://pay.chargily.net/api/v2"
	} else {
		return nil, errors.New("Chargily api mode must be either 'live' or 'test'.")
	}

	return &Chargily{
		baseUrl: apiMode,
		apiKey:  apiKey,
	}, nil
}

// Create customer method
func (c *Chargily) CreateCustomer(customerRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(customerRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customers", c.baseUrl), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Update customer method
func (c *Chargily) UpdateCustomer(id string, customerRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(customerRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/%s", c.baseUrl, id), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve customer method
func (c *Chargily) RetrieveCustomer(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Retrieve a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Retrieve a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Retrieve a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// List customers method
func (c *Chargily) ListCustomers(quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/customers?per_page=%d", c.baseUrl, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not List customers, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not List customers, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not List customers, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Delete a customer method
func (c *Chargily) DeleteCustomer(id string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/customers/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Delete a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Delete a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not Delete a customer, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Create a product method
func (c *Chargily) CreateProduct(productRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(productRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/products", c.baseUrl), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Update a product method
func (c *Chargily) UpdateProduct(id string, productRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(productRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/products/%s", c.baseUrl, id), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve a product method
func (c *Chargily) RetrieveProduct(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// List products method
func (c *Chargily) ListProducts(quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products?per_page=%d", c.baseUrl, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list products, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list products, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list products, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Delete a product method
func (c *Chargily) DeleteProduct(id string) ([]byte, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/products/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not delete a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not delete a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not delete a product, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve a product's price method
func (c *Chargily) RetrieveProductPrice(id string, quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/products/%s/prices?per_page=%d", c.baseUrl, id, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product's price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product's price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a product's price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Create a price method
func (c *Chargily) CreatePrice(priceRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(priceRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/prices", c.baseUrl), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Update a price method
func (c *Chargily) UpdatePrice(id string, priceMetadata interface{}) ([]byte, error) {
	requestData, err := json.Marshal(priceMetadata)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/prices/%s", c.baseUrl, id), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve a price method
func (c *Chargily) RetrievePrice(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/prices/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a price, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// List prices method
func (c *Chargily) ListPrices(quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/prices?per_page=%d", c.baseUrl, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list prices, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list prices, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list prices, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Create checkout method
func (c *Chargily) CreateCheckout(checkoutRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(checkoutRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/checkouts", c.baseUrl), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not create checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve checkout method
func (c *Chargily) RetrieveCheckout(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/checkouts/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// List checkouts method
func (c *Chargily) ListCheckouts(quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/checkouts?per_page=%d", c.baseUrl, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list checkouts, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list checkouts, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list checkouts, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve checkout's items method
func (c *Chargily) RetrieveCheckoutItems(id string, quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/checkouts/%s/items?per_page=%d", c.baseUrl, id, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout's items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout's items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve checkout's items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Expire a checkout method
func (c *Chargily) ExpireCheckout(id string) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/checkouts/%s/expire", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not expire a checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not expire a checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not expire a checkout, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Create a payment link method
func (c *Chargily) CreatePaymentLink(paymentLinkRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(paymentLinkRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/payment-links", c.baseUrl), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not create a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Update a payment link method
func (c *Chargily) UpdatePaymentLink(id string, paymentLinkRequest interface{}) ([]byte, error) {
	requestData, err := json.Marshal(paymentLinkRequest)
	if err != nil {
		formattedError := fmt.Sprintf("Error parsing request: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/payment-links/%s", c.baseUrl, id), bytes.NewBuffer(requestData))
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not update a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve a payment link method
func (c *Chargily) RetrievePaymentLink(id string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment-links/%s", c.baseUrl, id), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// List payment links method
func (c *Chargily) ListPaymentLinks(quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment-links?per_page=%d", c.baseUrl, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list payment links, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list payment links, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not list payment links, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}

// Retrieve a payment link items method
func (c *Chargily) RetrievePaymentLinkItems(id string, quantity int) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/payment-links/%s/items?per_page=%d", c.baseUrl, id, quantity), nil)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		formattedError := fmt.Sprintf("Could not retrieve a payment link items, error: %s", err.Error())
		return nil, errors.New(formattedError)
	}

	return respBody, nil
}
