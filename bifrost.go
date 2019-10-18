package bifrost

import (
	"bifrost/providers/registry"
	"bifrost/providers/types"
	"fmt"
)

// DecryptSecretFromStorage downloads the file from Storage to memory and decrypts it, returning the secret
func DecryptSecretFromStorage(provider types.ProviderConst, bucket string, path string, key string) ([]byte, error) {
	data, err := registry.ProviderRegistry[string(provider)].Download(bucket, path)
	if err != nil {
		return nil, fmt.Errorf("failed to download, %v", err)
	}
	secret, err := DecryptSecret(provider, key, data)

	return secret, err
}

// DecryptSecretFromStorageAsString downloads the file from Storage to memory and decrypts it, returning the secret as a string
func DecryptSecretFromStorageAsString(provider types.ProviderConst, bucket string, path string, key string) (string, error) {
	secret, err := DecryptSecretFromStorage(provider, bucket, path, key)
	return string(secret), err
}

// DecryptSecret decrypts the data returning the secret
func DecryptSecret(provider types.ProviderConst, key string, data []byte) ([]byte, error) {
	secret, err := registry.ProviderRegistry[string(provider)].Decrypt(key, data)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt, %v", err)
	}

	return secret, err
}

// DecryptSecretAsString decrypts the data, returning the secret as a string
func DecryptSecretAsString(provider types.ProviderConst, key string, data []byte) (string, error) {
	secret, err := DecryptSecret(provider, key, data)
	return string(secret), err
}

// DecryptSecretFromString decrypts the data from a provided string, returning the secret as a string
func DecryptSecretFromString(provider types.ProviderConst, key string, data string) ([]byte, error) {
	secret, err := DecryptSecret(provider, key, []byte(data))
	return secret, err
}

// DecryptSecretFromStringAsString decrypts the data from a provided string, returning the secret as a string
func DecryptSecretFromStringAsString(provider types.ProviderConst, key string, data string) (string, error) {
	secret, err := DecryptSecretAsString(provider, key, []byte(data))
	return secret, err
}

// EncryptSecretToStorage encrypts the data, saving it to the provided bucket and path on your provider.
func EncryptSecretToStorage(provider types.ProviderConst, bucket string, path string, key string, data []byte) error {
	ciphertext, err := EncryptSecret(provider, key, data)
	success, err := registry.ProviderRegistry[string(provider)].Upload(bucket, path, ciphertext)
	if err != nil || !success {
		return fmt.Errorf("failed to upload to storage, %v", err)
	}

	return nil
} 

// EncryptSecretToStorageFromString encrypts the data from a string, saving it to the provided bucket and path on your provider.
func EncryptSecretToStorageFromString(provider types.ProviderConst, bucket string, path string, key string, data string) error {
	return EncryptSecretToStorage(provider, bucket, path, key, []byte(data))
} 

// EncryptSecret encrypts the data, returning the bytes.
func EncryptSecret(provider types.ProviderConst, key string, data []byte) ([]byte, error) {
	ciphertext, err := registry.ProviderRegistry[string(provider)].Encrypt(key, data)
	if err != nil {
		return nil, fmt.Errorf("failed encrypt, %v", err)
	}

	return ciphertext, nil
} 

// EncryptSecretAsString encrypts the data from a bytes, returning a string.
func EncryptSecretAsString(provider types.ProviderConst, key string, data []byte) (string, error) {
	ciphertext, err := EncryptSecret(provider, key, data)
	return string(ciphertext), err
} 

// EncryptSecretFromString encrypts the data from a string, returning byte data.
func EncryptSecretFromString(provider types.ProviderConst, key string, data string) ([]byte, error) {
	ciphertext, err := EncryptSecret(provider, key, []byte(data))
	return ciphertext, err
} 

// EncryptSecretFromStringAsString encrypts the data from a string, returning string data..
func EncryptSecretFromStringAsString(provider types.ProviderConst, key string, data string) (string, error) {
	ciphertext, err := EncryptSecretAsString(provider, key, []byte(data))
	return ciphertext, err
} 