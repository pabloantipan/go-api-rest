package datastore

import (
	"context"
	"encoding/json"
	"learning/config"
	"log"
	"os"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"
)

type ServiceAccount struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func NewDatastoreClient(cfg *config.Config) *datastore.Client {
	// serviceAccountPath := os.Getenv("DATASTORE_SERVICE_ACCOUNT_PATH")
	// log.Print("oli", cfg.DatastoreServiceAccountPath)
	if cfg.DatastoreServiceAccountPath == "" {
		log.Printf("DATASTORE_SERVICE_ACCOUNT_PATH environment variable not set")
		return nil
	}

	log.Print("DATASTORE_SERVICE_ACCOUNT_PATH", cfg.DatastoreServiceAccountPath)

	data, err := os.ReadFile(cfg.DatastoreServiceAccountPath)
	if err != nil {
		log.Printf("Failed to read service account file: %v", err)
		return nil
	}

	var sa ServiceAccount
	if err := json.Unmarshal(data, &sa); err != nil {
		log.Printf("Failed to unmarshal service account file: %v", err)
		return nil
	}

	ctx := context.Background()
	client, err := datastore.NewClient(
		ctx,
		sa.ProjectID,
		option.WithCredentialsFile(cfg.DatastoreServiceAccountPath),
	)

	if err != nil {
		log.Fatalf("Failed to create datastore client: %v", err)
		return nil
	}

	return client
}
