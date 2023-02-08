package internal

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

	"github.com/spf13/viper"
)

func GetBlobServiceClient() (*azblob.Client, error) {

	// Create a default Azure credential
	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	// Create blob client
	return azblob.NewClient(fmt.Sprintf("https://%s.blob.core.windows.net", viper.GetString("storageAccountName")), credential, nil)
}
