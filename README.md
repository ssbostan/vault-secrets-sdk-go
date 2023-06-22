# vault-secrets-sdk-go

**HCP Vault Secrets** SDK for Go programming language.

HCP Vault Secrets is a simplified KV engine that allows you to store your application's secrets and credentials in a centralized and secure way. By using HCP Vault Secrets, you don't need to deal with the HashiCorp Vault cluster! You need to create an application, a bunch of secrets, go.

Read [more](https://developer.hashicorp.com/vault/tutorials/hcp-vault-secrets-get-started) on HashiCorp Developer Portal.

## Quick start guide:

```go
package main

import (
	"fmt"
	"log"
	"strings"

	vaultsecrets "github.com/ssbostan/vault-secrets-sdk-go"
)

func main() {
	client, err := vaultsecrets.NewClient(
		"HCP_ORGANIZATION_ID",
		"HCP_PROJECT_ID",
		"HCP_APPLICATION_NAME",
		"HCP_CLIENT_ID",
		"HCP_CLIENT_SECRET",
	)
	if err != nil {
		log.Fatalln(err)
	}

	secret, err := client.Get("MY_APP_SECRET")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Get method for MY_APP_SECRET: %s\n", secret)

	secretsList, err := client.List()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("List method for listing secrets: %s\n", strings.Join(secretsList, ", "))

	secrets, err := client.GetAll()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("GetAll method to get all secrets:")
	for key, value := range secrets {
		fmt.Printf("%s: %s\n", key, value)
	}
}
```

## Available functions/methods:

  - **NewClient(o, p, a, clientID, clientSecret)**: Returns ready to use Client struct.
  - **Client.Authenticate()**: Authenticates and gets API Access Token.
  - **Client.Get(name)**: Fetches secret value from the specified application.
  - **Client.List()**: List all secrets in the specified application.
  - **Client.GetAll()**: Get value of all secrets in the specified application.

Copyright 2023 Saeid Bostandoust <ssbostan@yahoo.com>
