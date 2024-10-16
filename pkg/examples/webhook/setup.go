package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/Chargily/chargily-pay-go/pkg/models"
)


func SetupWebhook(client * chargily.Client){
	// Setup the webhook endpoint to listen for webhook events on `/webhook`
	client.Webhook.SetupHandler("/webhook", handleEvent)
	// Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// Custom event handler function that processes events
func handleEvent(eventType string, event models.WebhookEvent) {
	// Handle different event types (e.g., `checkout.paid`, etc.)
	switch eventType {
	case "checkout.paid":
		// Process a successful payment
		fmt.Printf("Received paid event: %+v\n", event)
		// Do something with the event data (e.g., update order status)
	case "checkout.failed":
		// Handle failed payment
		fmt.Println("Payment failed for event ID:", event.ID)
	default:
		fmt.Printf("Unhandled event type: %s\n", eventType)
	}
}