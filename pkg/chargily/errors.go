package chargily

import "errors"

// Custom error for invalid mode
var ErrInvalidMode = errors.New("invalid mode: must be 'prod' or 'test'")
