package vault

import (
	"fmt"
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

func GetValue(client *api.Client) {
	// Define the secret's path
	// Note: For KV v2, the path is usually prefixed with "secret/data/"
	secretPath := "ick-test/data/ick-ips/config-service/workflow-service/dev1"

	// Read the secret
	secret, err := client.Logical().Read(secretPath)
	if err != nil {
		log.Fatalf("Failed to read secret: %s", err)
	}
	if secret == nil {
		log.Fatalf("Secret not found at path: %s", secretPath)
	}

	// For KV v2, the secret data is under the "data" key
	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		log.Fatalf("Data type assertion failed for secret: %s", secretPath)
	}

	// Assuming the secret has a key "password"
	password, ok := data["KAFKA_ENABLED"].(string)
	if !ok {
		log.Fatalf("KAFKA_ENABLED not found in secret: %s", secretPath)
	}

	fmt.Printf("Retrieved KAFKA_ENABLED: %s\n", password)
}
