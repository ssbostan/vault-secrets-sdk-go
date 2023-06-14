package vaultsecrets

const (
	AuthURL           string = "https://auth.hashicorp.com/oauth/token"
	AuthTokenAudience string = "https://api.hashicorp.cloud"
	AuthGrantType     string = "client_credentials"
	VaultSecretsURL   string = "https://api.cloud.hashicorp.com/secrets/2023-06-13"
)

type Client struct {
	OrganizationID  string
	ProjectID       string
	ApplicationName string
	ClientID        string
	ClientSecret    string
	AccessToken     string
}
