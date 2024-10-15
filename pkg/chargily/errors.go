package chargily

import "errors"

// GeneralError represents the structure of a generic error response.
type GeneralError struct {
    Message string               `json:"message"` // A general error message.
    Errors  map[string][]string  `json:"errors"`  // A map where the key is the field name and the value is a slice of error messages.
}


// Custom error for invalid mode
var ErrInvalidMode = errors.New("invalid mode: must be 'prod' or 'test'")
