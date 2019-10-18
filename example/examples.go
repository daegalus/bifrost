package example

import (
	"bifrost"
	"bifrost/bfconsts"
	"bifrost/providers/registry"
	"encoding/base64"
	"fmt"
)

func example() {
  registry.LoadProviders()
	dataStr := "The quick brown fox jumped over the dog."
	dataBytes := []byte(dataStr)

	bucket := "some-bucket"
	path := "yulian/bifrost-test"
	kms := "projects/some-project/locations/global/keyRings/some-keyring/cryptoKeys/some-key"
	
	fmt.Println("Encrypting...")
	encDataBytes, _ := bifrost.EncryptSecret(bfconsts.ProviderGCP, kms, dataBytes)
	fmt.Println("Encrypting as string...")
	encDataStr, _ := bifrost.EncryptSecretAsString(bfconsts.ProviderGCP, kms, dataBytes)
	fmt.Println("Encrypting from string...")
	encDataBytesFS, _ := bifrost.EncryptSecretFromString(bfconsts.ProviderGCP, kms, dataStr)
	fmt.Println("Encrypting from string as string...")
	encDataStrFS, _ := bifrost.EncryptSecretFromStringAsString(bfconsts.ProviderGCP, kms, dataStr)
	fmt.Println("Encrypting to storage...")
	bifrost.EncryptSecretToStorage(bfconsts.ProviderGCP, bucket, path, kms, dataBytes)
	fmt.Println("Encrypting to storage from string...")
	bifrost.EncryptSecretToStorageFromString(bfconsts.ProviderGCP, bucket, path, kms, dataStr)

	fmt.Println("Decrypting...")
	decDataBytes, err := bifrost.DecryptSecret(bfconsts.ProviderGCP, kms, encDataBytes) 
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypting as string...")
	decDataStr, _ := bifrost.DecryptSecretAsString(bfconsts.ProviderGCP, kms, encDataBytesFS)
	fmt.Println("Decrypting from string...")
	decDataBytesFS, _ := bifrost.DecryptSecretFromString(bfconsts.ProviderGCP, kms, encDataStr)
	fmt.Println("Decrypting from string as string...")
	decDataStrFS, _ := bifrost.DecryptSecretFromStringAsString(bfconsts.ProviderGCP, kms, encDataStrFS)
	fmt.Println("Decrypting from storage...")
	decDataBytesStor, errS := bifrost.DecryptSecretFromStorage(bfconsts.ProviderGCP, bucket, path, kms)
	if errS != nil {
		panic(errS)
	}
	fmt.Println("Decrypting from storage as string...")
	decDataStrStor, _ := bifrost.DecryptSecretFromStorageAsString(bfconsts.ProviderGCP, bucket, path, kms)

	fmt.Printf("Encrypted Bytes: %v\n", encDataBytes)
	fmt.Printf("Encrypted String: %s\n", base64.StdEncoding.EncodeToString([]byte(encDataStr)))
	fmt.Printf("Encrypted Bytes From String: %v\n", encDataBytesFS)
	fmt.Printf("Encrypted String From String: %s\n", base64.StdEncoding.EncodeToString([]byte(encDataStrFS)))

	fmt.Printf("Decrypted Bytes: %v\n", decDataBytes)
	fmt.Printf("Decrypted String: %s\n", decDataStr)
	fmt.Printf("Decrypted Bytes From String: %v\n", decDataBytesFS)
	fmt.Printf("Decrypted String From String: %s\n", decDataStrFS)
	fmt.Printf("Decrypted Bytes From Storage: %v\n", decDataBytesStor)
	fmt.Printf("Decrypted String From Storage: %s\n", decDataStrStor)
}