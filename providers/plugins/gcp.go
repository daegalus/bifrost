package plugins

import (
	"bifrost/providers/types"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"

	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/storage"
)

// GCP struct for the GCP provider to store the clients to the GCP API
type GCP struct {
	KMSClient *kms.KeyManagementClient
	KMSContext *context.Context
	StorageClient *storage.Client
	StorageContext *context.Context
}

// Setup creates the clients for use to talk to GCP.
func (gcp GCP) Setup() {
	ctx := context.Background()
	storagec, err := storage.NewClient(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to create GCP Storage client. %v", err))
	}
	kmsc, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to create GCP KMS client. %v", err))
	}
	gcp.KMSClient = kmsc
	gcp.KMSContext = &ctx
	gcp.StorageClient = storagec
	gcp.StorageContext = &ctx
}

// Encrypt encrypts the data using GCP KMS
func (gcp GCP) Encrypt(key string, data []byte) ([]byte, error) {
	req := &kmspb.EncryptRequest{
		Name: key,
		Plaintext: data,
	}
	resp, err := gcp.KMSClient.Encrypt(*gcp.KMSContext, req)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt data, %v", err)
	}

  return resp.GetCiphertext(), nil
}

// Decrypt decrypts the data using GCP KMS
func (gcp GCP) Decrypt(key string, data []byte) ([]byte, error) {
	req := &kmspb.DecryptRequest{
		Name: key,
		Ciphertext: data,
	}
	resp, err := gcp.KMSClient.Decrypt(*gcp.KMSContext, req)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data, %v", err)
	}
	
	return resp.GetPlaintext(), nil
}

// Download downloads the data from a GCP Cloud Storage Bucket
func (gcp GCP) Download(bucket string, path string) ([]byte, error) {
	bkt := gcp.StorageClient.Bucket(bucket)
	obj := bkt.Object(path)

	var data bytes.Buffer
	buf := bufio.NewWriter(&data)

	r, err := obj.NewReader(*gcp.StorageContext) 
	if err != nil {
			return nil, fmt.Errorf("failed to create object reader, %v", err)
	}
	defer r.Close()
	if _, err := io.Copy(buf, r); err != nil {
		return nil, fmt.Errorf("failed to write data to buffer, %v", err)
	}
	return data.Bytes(), nil
}

// Upload uploads the data to a GCP Cloud Storage Bucket
func (gcp GCP) Upload(bucket string, path string, data []byte) (bool, error) {
	bkt := gcp.StorageClient.Bucket(bucket)
	obj := bkt.Object(path)
	
	w := obj.NewWriter(*gcp.StorageContext)
	if _, err := w.Write(data); err != nil {
			return false, fmt.Errorf("failed to write data to GCS path, %v", err)
	}
	
	if err := w.Close(); err != nil {
			return false, fmt.Errorf("failed to close writer handle, %v", err)
	}

	return true, nil
}

// ProviderID the string id of this provider for lookup.
func (gcp GCP) ProviderID() string {
	return "gcp"
}

// Register self-registration with the ProviderRegistry
func (gcp GCP) Register(registry *map[string]types.Provider) {
	gcp.Setup()
	(*registry)[gcp.ProviderID()] = gcp 
}