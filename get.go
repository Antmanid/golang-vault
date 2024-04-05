package get

import (
	"fmt"
	"log"

	"github.com/hashicorp/vault/api"
)

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
