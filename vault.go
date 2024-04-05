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

func GetValue(client *api.Client, secretPath string) map[string]interface{} {

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

	return data
}

func WriteValue(client *api.Client, dstPath string, inputData map[string]interface{}) {

	// Write the secret
	output, err := client.Logical().Write(dstPath, inputData)
	if err != nil {
		log.Fatalf("failed to write secret: %v", err)
	}
	fmt.Println(output)

}
