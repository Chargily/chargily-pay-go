package chargily

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Chargily/chargily-pay-go/pkg/models"
)

type Webhook struct {
	client * Client
}

type EventHandler func(eventType string ,event models.WebhookEvent)

//wh : webhook
func (wh * Webhook) SetupHandler(path string, handler EventHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		// Extract signature
		signature := r.Header.Get("signature")

		// Check if signature is present
		if signature == "" {
			http.Error(w, "Missing signature", http.StatusBadRequest)
			return
		}

		// Read payload
		payload, err := io.ReadAll(r.Body)
		// Close the request body to free up resources
		defer r.Body.Close()
		if err != nil {
			http.Error(w, "Failed to read payload", http.StatusInternalServerError)
			return
		}


		// Verify signature
		err = wh.VerifySignature(payload, signature)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Parse JSON payload
		var event models.WebhookEvent
		if err := json.Unmarshal(payload, &event); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusInternalServerError)
			return
		}

		// Identify event type and call handler
  		eventType := string(event.Type)

		// Call user-defined handler
		handler(eventType, event)

		// Respond with 200 OK
		w.WriteHeader(http.StatusOK)
	})
}



// Reuseable signature verifier function 
func (wh * Webhook) VerifySignature(payload []byte, signature string) error{
	//compute the HMAC signature
	computedSignature := computeHMAC(payload, wh.client.apiKey)
	// Compare the computed signature with the received signature
	if !hmac.Equal([]byte(computedSignature), []byte(signature)) {
		// If they don't match, return an error indicating invalid signature
		return errors.New("invalid signature")
	}
	// If they match, return nil indicating valid signature
	return nil; 
}


// Helper function to compute HMAC
func computeHMAC(data []byte, key string) string {
	// Create a new HMAC hash using SHA256 and the provided API key
	h := hmac.New(sha256.New, []byte(key))

	// Write the data to the hasher and compute the final HMAC
	h.Write(data)

	// Return the computed HMAC as a hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}


