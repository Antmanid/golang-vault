package auth

import (
	"log"

	"github.com/hashicorp/vault/api"
)

// AddIntNum is to add two integer numbers
func AddIntNum(num1, num2 int) int {
	return num1 + num2
}

func TokenVaultClient(vaultAddr string, vaultToken string) *api.Client {

	// Initialize a Vault API client
	config := api.DefaultConfig()
	config.Address = vaultAddr
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("failed to create Vault client: %v", err)
	}
	client.SetToken(vaultToken)

	return client
}
