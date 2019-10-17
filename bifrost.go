package main

import (
	"bifrost/providers/registry"
	"bifrost/providers/types"
	"fmt"
)
func main() {}

const (
	// ProviderGCP constant for GCP
	ProviderGCP = types.ProviderConst("gcp")
	// ProviderAWS constant for AWS
	ProviderAWS = types.ProviderConst("aws")
)

// DecryptSecretFromStorage downloads the file from Storage to memory and decrypts it, returning the secret
func DecryptSecretFromStorage(provider types.ProviderConst, bucket string, path string, key string) ([]byte, error) {
	data, err := registry.ProviderRegistry[string(provider)].Download(bucket, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download, %v", err)
	}
	secret, err := registry.ProviderRegistry[string(provider)].Decrypt(key, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt, %v", err)
	}

	return secret, nil
}

// DecryptSecretFromStorageAsString downloads the file from Storage to memory and decrypts it, returning the secret as a string
func DecryptSecretFromStorageAsString(provider types.ProviderConst, bucket string, path string, key string) (string, error) {
	secret, err := DecryptSecretFromStorage(provider, bucket, path, key)
	return string(secret), err
}

// DecryptSecret decrypts the data returning the secret
func DecryptSecret(provider types.ProviderConst, data []byte, key string) ([]byte, error) {
	secret, err := registry.ProviderRegistry[string(provider)].Decrypt(key, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt, %v", err)
	}

	return secret, nil
}

// DecryptSecretAsString decrypts the data, returning the secret as a string
func DecryptSecretAsString(provider types.ProviderConst, data []byte, key string) (string, error) {
	secret, err := DecryptSecret(provider, data, key)
	return string(secret), err
}