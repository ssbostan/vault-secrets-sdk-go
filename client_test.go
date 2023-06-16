package vaultsecrets

import (
	"os"
	"testing"
)

var client = Client{
	OrganizationID:  os.Getenv("HCP_ORGANIZATION_ID"),
	ProjectID:       os.Getenv("HCP_PROJECT_ID"),
	ApplicationName: os.Getenv("HCP_APPLICATION_NAME"),
	ClientID:        os.Getenv("HCP_CLIENT_ID"),
	ClientSecret:    os.Getenv("HCP_CLIENT_SECRET"),
}

func TestAuthenticate(t *testing.T) {
	if err := client.Authenticate(); err != nil {
		t.Errorf("client Authenticate error: %s", err)
	}

	if client.AccessToken == "" {
		t.Errorf("client Authenticate error: Could not retrieve access token.")
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		secretName    string
		expectedValue string
	}{
		{"NOT_FOUND_SECRET", ""},
		{"SDK_TEST", "OK"},
	}

	for _, test := range tests {
		secret, _ := client.Get(test.secretName)
		if secret != test.expectedValue {
			t.Errorf("client Get error: Could not meet expected value.")
		}
	}
}

func TestNewClient(t *testing.T) {
	_, err := NewClient(
		os.Getenv("HCP_ORGANIZATION_ID"),
		os.Getenv("HCP_PROJECT_ID"),
		os.Getenv("HCP_APPLICATION_NAME"),
		os.Getenv("HCP_CLIENT_ID"),
		os.Getenv("HCP_CLIENT_SECRET"),
	)
	if err != nil {
		t.Errorf("client NewClient error: %s", err)
	}
}
