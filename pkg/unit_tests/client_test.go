package unit_tests

import (
	"testing"

	"github.com/Chargily/chargily-pay-go/pkg/chargily"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock for the rs type
type MockRequestSender struct {
	mock.Mock
}

func TestNewClient(t *testing.T) {
	tests := []struct {
		name        string
		apiKey      string
		mode        string
		expectedErr error
	}{
		{"Valid Prod Client", "test-api-key", "prod", nil},
		{"Valid Test Client", "test-api-key", "test", nil},
		{"Invalid Mode", "test-api-key", "invalid", chargily.ErrInvalidMode},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := chargily.NewClient(tt.apiKey, tt.mode)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}
