package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-test-api-key")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if apiKey != "my-test-api-key" {
		t.Errorf("expected API key %q, got %q", "my-test-api-key", apiKey)
	}
}

func TestGetAPIKeyNoAuthorizationHeader(t *testing.T) {
	headers := http.Header{}

	apiKey, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
	}

	if apiKey != "" {
		t.Errorf("expected empty API key, got %q", apiKey)
	}
}

func TestGetAPIKeyMalformedAuthorizationHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-test-api-key")

	apiKey, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if apiKey != "" {
		t.Errorf("expected empty API key, got %q", apiKey)
	}
}
