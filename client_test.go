package vaultsecrets

import (
	"os"
	"reflect"
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

func TestGetLatestVersion(t *testing.T) {
	tests := []struct {
		secretName    string
		expectedValue int
	}{
		{"NOT_FOUND_SECRET", 0},
		{"SDK_TEST", 1},
	}

	for _, test := range tests {
		version, _ := client.GetLatestVersion(test.secretName)
		if version != test.expectedValue {
			t.Errorf("client GetLatestVersion error: Could not meet expected value.")
		}
	}
}

func TestCreate(t *testing.T) {
	version, _ := client.GetLatestVersion("MY_APP_SECRET")
	newVersion, _ := client.Create("MY_APP_SECRET", "SDK_TEST")
	if version+1 != newVersion {
		t.Errorf("client Create error: New version of secret was not created")
	}
}

func TestList(t *testing.T) {
	secrets, err := client.List()
	if err != nil {
		t.Errorf("client List error: %s", err)
	}

	if !reflect.DeepEqual(secrets, []string{"MY_APP_SECRET", "SDK_TEST"}) {
		t.Errorf("client List error: Could not meet expected value.")
	}
}

func TestGetAll(t *testing.T) {
	secrets, err := client.GetAll()
	if err != nil {
		t.Errorf("client GetAll error: %s", err)
	}

	if !reflect.DeepEqual(
		secrets,
		map[string]string{
			"MY_APP_SECRET": "TEST",
			"SDK_TEST":      "OK",
		},
	) {
		t.Errorf("client GetAll error: Could not meet expected value.")
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
