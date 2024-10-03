package chargilypaygo

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	CUSTOMER_ID     string
	PRODUCT_ID      string
	PRICE_ID        string
	CHECKOUT_ID     string
	PAYMENT_LINK_ID string
)

func TestCreateCustomer(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	customerReq := map[string]any{
		"name": "testUser",
	}

	resp, err := client.CreateCustomer(customerReq)
	if err != nil {
		t.Errorf("Error creating customer: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when creating customer: %s", err.Error())
		return
	}

	if data["name"] != "testUser" {
		t.Errorf("Expected 'testUser' as a result when creating customer, but got: %s", string(resp))
		return
	}

	if id, ok := data["id"].(string); ok {
		CUSTOMER_ID = id
	} else {
		t.Error("Could not convert data['id'] to a string when creating customer")
		return
	}
}

func TestUpdateCustomer(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	customerReq := map[string]any{
		"name": "anotherTestUser",
	}

	resp, err := client.UpdateCustomer(CUSTOMER_ID, customerReq)
	if err != nil {
		t.Errorf("Error updating customer: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when updating customer: %s", err.Error())
		return
	}

	if data["name"] != "anotherTestUser" {
		t.Errorf("Expected 'anotherTestUser' as a result when updating customer, but got: %s", data["name"])
		return
	}
}

func TestRetrieveCustomer(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrieveCustomer(CUSTOMER_ID)
	if err != nil {
		t.Errorf("Error retrieving customer: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving customer: %s", err.Error())
		return
	}

	if data["id"] != CUSTOMER_ID {
		t.Errorf("Expected %s as a result when retrieving customer, but got: %s", CUSTOMER_ID, data["id"])
		return
	}
}

func TestListingCustomer(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ListCustomers(1)
	if err != nil {
		t.Errorf("Error listing customers: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when listing customers: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error listing customers, expected an array of key-value pairs,
		but got: %v`, string(resp))
		return
	}
}

func TestDeleteCustomer(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.DeleteCustomer(CUSTOMER_ID)
	if err != nil {
		t.Errorf("Error deleting customer: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when deleting customer: %s", err.Error())
		return
	}

	if data["id"] != CUSTOMER_ID {
		t.Errorf("Expected %s as a result when deleting customer, but got: %s", CUSTOMER_ID, data["id"])
		return
	}
}

func TestCreateProduct(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	productReq := map[string]any{
		"name": "testProduct",
	}

	resp, err := client.CreateProduct(productReq)
	if err != nil {
		t.Errorf("Error creating product: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when creating a product: %s", err.Error())
		return
	}

	if data["name"] != "testProduct" {
		t.Errorf("Expected 'testProduct' as a result when creating a product, but got: %s", data["name"])
		return
	}

	if id, ok := data["id"].(string); ok {
		PRODUCT_ID = id
	} else {
		t.Errorf("Could not convert data['id'] to a string when creating a product")
	}
}

func TestUpdateProduct(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	productReq := map[string]any{
		"name": "anotherTestProduct",
	}

	resp, err := client.UpdateProduct(PRODUCT_ID, productReq)
	if err != nil {
		t.Errorf("Error updating product: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when updating a product: %s", err.Error())
		return
	}

	if data["name"] != "anotherTestProduct" {
		t.Errorf("Expected 'anotherTestProduct' as a result when updating a product, but got: %s", data["name"])
		return
	}
}

func TestRetrieveProduct(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrieveProduct(PRODUCT_ID)
	if err != nil {
		t.Errorf("Error retrieving product: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving a product: %s", err.Error())
		return
	}

	if data["id"] != PRODUCT_ID {
		t.Errorf("Expected '%s' as a result when retrieving a product, but got: %s", PRODUCT_ID, data["id"])
		return
	}
}

func TestListProducts(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ListProducts(1)
	if err != nil {
		t.Errorf("Error listing products: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when listing products: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error listing products, expected an array of key-value pairs,
		but got: %v`, string(resp))
		return
	}
}

func TestRetrieveProductPrices(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrieveProductPrice(PRODUCT_ID, 1)
	if err != nil {
		t.Errorf("Error retrieving product's prices: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving product's prices: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error retrieving product's prices, expected an array of key-value pairs,
		but got: %v`, string(resp))
		return
	}
}

func TestCreatePrice(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	priceReq := map[string]any{
		"amount":     73.26,
		"currency":   "dzd",
		"product_id": PRODUCT_ID,
	}

	resp, err := client.CreatePrice(priceReq)
	if err != nil {
		t.Errorf("Error creating price: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when creating price: %s", err.Error())
		return
	}

	if data["amount"] != 73.26 {
		t.Errorf("Expected '%f' as a result when creating price, but got: %v", 73.26, string(resp))
		return
	}

	if id, ok := data["id"].(string); ok {
		PRICE_ID = id
	} else {
		t.Errorf("Could not convert data['id'] to a string when creating price")
	}
}

func TestUpdatePrice(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	metadataReq := map[string]any{
		"metadata": []map[string]any{
			{
				"testKey": "testValue",
			},
		},
	}

	resp, err := client.UpdatePrice(PRICE_ID, metadataReq)
	if err != nil {
		t.Errorf("Error updating price: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when updating price: %s", err.Error())
		return
	}

	// Checking if the 'metadata' field exists in the response
	// and whether it's an array of key-value pairs
	// and check if the key-value pairs match what we created
	// if not we throw an error.
	if metadata, ok := data["metadata"].([]interface{}); ok {
		if metadataMap, ok := metadata[0].(map[string]any); ok {
			if metadataMap["testKey"] != "testValue" {
				t.Errorf("Expected 'testValue' as a result when updating price, but got: %s", string(resp))
				return
			}
		} else {
			t.Errorf("could not convert metadata[0] to a map when updating price")
		}
	} else {
		t.Errorf("data['metadata'] is not an array of key-value pairs when updating price, or it doesn't exist")
	}
}

func TestRetrievePrice(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrievePrice(PRICE_ID)
	if err != nil {
		t.Errorf("Error retrieving price: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving price: %s", err.Error())
		return
	}

	if data["id"] != PRICE_ID {
		t.Errorf("Expected '%s' as a result when retrieving price, but got: %s", PRICE_ID, string(resp))
		return
	}
}

func TestListPrices(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ListPrices(1)
	if err != nil {
		t.Errorf("Error listing prices: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when listing prices: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error listing prices, expected an array of key-value pairs,
		but got: %s`, string(resp))
		return
	}
}

func TestCreateCheckout(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	checkoutReq := map[string]any{
		"items": []map[string]any{
			{
				"price":    PRICE_ID,
				"quantity": 1,
			},
		},
		"success_url": "https://test.com",
	}

	resp, err := client.CreateCheckout(checkoutReq)
	if err != nil {
		t.Errorf("Error creating checkout: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when creating checkout: %s", err.Error())
		return
	}

	if data["success_url"] != "https://test.com" {
		t.Errorf("Expected 'https://test.com' as a result when creating a checkout, but got: %s", string(resp))
		return
	}

	if id, ok := data["id"].(string); ok {
		CHECKOUT_ID = id
	} else {
		t.Errorf("Could not convert data['id'] to a string when creating checkout")
	}
}

func TestRetrieveCheckout(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrieveCheckout(CHECKOUT_ID)
	if err != nil {
		t.Errorf("Error retrieving checkout: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving checkout: %s", err.Error())
		return
	}

	if data["id"] != CHECKOUT_ID {
		t.Errorf("Expected '%s' as a result when retrieving checkout, but got: %s", CHECKOUT_ID, string(resp))
		return
	}
}

func TestListCheckouts(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ListCheckouts(1)
	if err != nil {
		t.Errorf("Error listing checkouts: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when listing checkouts: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error listing checkouts, expected an array of key-value pairs,
		but got: %s`, string(resp))
		return
	}
}

func TestRetrieveCheckoutItems(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrieveCheckoutItems(CHECKOUT_ID, 1)
	if err != nil {
		t.Errorf("Error retrieving checkout items: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving checkout items: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error retrieving checkout's items, expected an array of key-value pairs,
		but got: %s`, string(resp))
		return
	}
}

func TestExpireCheckout(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ExpireCheckout(CHECKOUT_ID)
	if err != nil {
		t.Errorf("Error expiring checkout items: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when expiring checkout: %s", err.Error())
		return
	}

	if data["id"] != CHECKOUT_ID {
		t.Errorf("Expected '%s' as a result when expiring checkout, but got: %s", CHECKOUT_ID, string(resp))
		return
	}
}

func TestCreatePaymentLink(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	paymentLinkReq := map[string]any{
		"name": "testPaymentLink",
		"items": []map[string]any{
			{
				"price":    PRICE_ID,
				"quantity": 1,
			},
		},
	}

	resp, err := client.CreatePaymentLink(paymentLinkReq)
	if err != nil {
		t.Errorf("Error creating payment link: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when creating payment link: %s", err.Error())
		return
	}

	if data["name"] != "testPaymentLink" {
		t.Errorf("Expected 'testPaymentLink' as a result when creating payment link, but got: %s", string(resp))
		return
	}

	if id, ok := data["id"].(string); ok {
		PAYMENT_LINK_ID = id
	} else {
		t.Errorf("Could not convert data['id'] to a string when creating a payment link")
	}
}

func TestUpdatePaymentLink(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	paymentLinkReq := map[string]any{
		"name": "anotherTestPaymentLink",
		"items": []map[string]any{
			{
				"price":    PRICE_ID,
				"quantity": 1,
			},
		},
	}

	resp, err := client.UpdatePaymentLink(PAYMENT_LINK_ID, paymentLinkReq)
	if err != nil {
		t.Errorf("Error updating payment link: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when updating payment link: %s", err.Error())
		return
	}

	if data["name"] != "anotherTestPaymentLink" {
		t.Errorf("Expected 'anotherTestPaymentLink' as a result when updating payment link, but got: %s", string(resp))
		return
	}
}

func TestRetrievePaymentLink(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrievePaymentLink(PAYMENT_LINK_ID)
	if err != nil {
		t.Errorf("Error retrieving payment link: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving payment link: %s", err.Error())
		return
	}

	if data["id"] != PAYMENT_LINK_ID {
		t.Errorf("Expected '%s' as a result when retrieving payment link, but got: %s", PAYMENT_LINK_ID, string(resp))
		return
	}
}

func TestListPaymentLinks(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.ListPaymentLinks(1)
	if err != nil {
		t.Errorf("Error listing payment links: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when listing payment links: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error listing payment links, expected an array of key-value pairs,
		but got: %s`, string(resp))
		return
	}
}

func TestRetrievePaymentLinkItems(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.RetrievePaymentLinkItems(PAYMENT_LINK_ID, 1)
	if err != nil {
		t.Errorf("Error retrieving payment link items: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when retrieving payment link items: %s", err.Error())
		return
	}

	// Checking if the 'data' field exists in the response
	// and whether it's an array
	// if not we throw an error.
	if _, ok := data["data"].([]interface{}); !ok {
		t.Errorf(`Error retrieving payment link items, expected an array of key-value pairs,
		but got: %s`, string(resp))
		return
	}
}

func TestDeleteProduct(t *testing.T) {
	client, err := NewChargily("test", os.Getenv("API_KEY"))
	if err != nil {
		t.Errorf("Error initializing chargily: %s", err.Error())
		return
	}

	resp, err := client.DeleteProduct(PRODUCT_ID)
	if err != nil {
		t.Errorf("Error deleting product: %s", err.Error())
		return
	}

	data := map[string]any{}

	err = json.Unmarshal(resp, &data)
	if err != nil {
		t.Errorf("Error unmarshaling response when deleting product: %s", err.Error())
		return
	}

	if data["id"] != PRODUCT_ID {
		t.Errorf("Expected '%s' as a result when deleting product, but got: %s", PRODUCT_ID, string(resp))
		return
	}
}
