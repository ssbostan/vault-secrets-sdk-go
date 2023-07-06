package vaultsecrets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Authenticate() error {
	authRequestData, err := json.Marshal(
		map[string]string{
			"audience":      AuthTokenAudience,
			"grant_type":    AuthGrantType,
			"client_id":     c.ClientID,
			"client_secret": c.ClientSecret,
		},
	)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequest("POST", AuthURL, bytes.NewReader(authRequestData))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"Authentication failed with %d status code.",
			httpResponse.StatusCode,
		)
	}

	httpResponseData, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	var authResponseData map[string]interface{}
	if err := json.Unmarshal(httpResponseData, &authResponseData); err != nil {
		return err
	}

	c.AccessToken = authResponseData["access_token"].(string)
	return nil
}

func (c *Client) Get(name string) (string, error) {
	httpRequest, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/organizations/%s/projects/%s/apps/%s/open/%s",
			VaultSecretsURL, c.OrganizationID, c.ProjectID, c.ApplicationName, name,
		),
		nil,
	)
	if err != nil {
		return "", err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return "", err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"Could not retrieve '%s' secret with %d status code.",
			name, httpResponse.StatusCode,
		)
	}

	httpResponseData, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	var secretData map[string]interface{}
	if err := json.Unmarshal(httpResponseData, &secretData); err != nil {
		return "", err
	}

	return secretData["secret"].(map[string]interface{})["version"].(map[string]interface{})["value"].(string), nil
}

func (c *Client) List() ([]string, error) {
	httpRequest, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/organizations/%s/projects/%s/apps/%s/secrets",
			VaultSecretsURL, c.OrganizationID, c.ProjectID, c.ApplicationName,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"Could not retrieve secrets with %d status code.", httpResponse.StatusCode,
		)
	}

	httpResponseData, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var secretsData map[string]interface{}
	if err := json.Unmarshal(httpResponseData, &secretsData); err != nil {
		return nil, err
	}

	var secrets []string

	for _, secret := range secretsData["secrets"].([]interface{}) {
		secrets = append(secrets, secret.(map[string]interface{})["name"].(string))
	}

	return secrets, nil
}

func (c *Client) GetAll() (map[string]string, error) {
	secrets, err := c.List()
	if err != nil {
		return nil, err
	}

	secretValues := make(map[string]string)

	for _, secret := range secrets {
		secretValues[secret], err = c.Get(secret)
		if err != nil {
			return nil, err
		}
	}

	return secretValues, nil
}

func NewClient(
	organizationID string,
	projectID string,
	applicationName string,
	clientID string,
	clientSecret string,
) (Client, error) {
	client := Client{
		OrganizationID:  organizationID,
		ProjectID:       projectID,
		ApplicationName: applicationName,
		ClientID:        clientID,
		ClientSecret:    clientSecret,
	}

	if err := client.Authenticate(); err != nil {
		return Client{}, err
	}

	return client, nil
}
