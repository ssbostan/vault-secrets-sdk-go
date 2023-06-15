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

	vaultsecrets "github.com/ssbostan/vault-secrets-sdk-go"
)

func main() {
	client := vaultsecrets.Client{
		OrganizationID:  "HCP_ORGANIZATION_ID",
		ProjectID:       "HCP_PROJECT_ID",
		ApplicationName: "HCP_APPLICATION_NAME",
		ClientID:        "HCP_CLIENT_ID",
		ClientSecret:    "HCP_CLIENT_SECRET",
	}
	if err := client.Authenticate(); err != nil {
		log.Fatal(err)
	}

	secret, err := client.Get("MY_APP_SECRET")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(secret)
}
```

## Available methods:

  - **Client.Authenticate()**: Authenticates and gets API Access Token.
  - **Client.Get(name)**: Fetches secret value from the specified application.

## TODO:

  - [ ] Implement TDD approach.
  - [ ] Implement NewClient function.
  - [ ] Implement List method.
  - [ ] Implement GetAll method.

Copyright 2023 Saeid Bostandoust <ssbostan@yahoo.com>
