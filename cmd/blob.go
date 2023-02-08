/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azure-storage-golang-demos/internal"
	"context"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// blobCmd represents the blob command
var blobCmd = &cobra.Command{
	Use:   "blob",
	Short: "Create a new blob",
	Long:  `It creates a new blob in your storage account.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		internal.CheckCredentials()
	},
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		// Get flags values
		containerName := cmd.Flag("container").Value.String()
		blobName := cmd.Flag("name").Value.String()
		path := cmd.Flag("path").Value.String()

		// Get blob service client
		serviceClient, err := internal.GetBlobServiceClient()

		if err != nil {
			log.Fatal("Error trying to get service client: " + err.Error())
		}

		// Try to open the file you want to upload
		file, err := os.Open(path)

		if err != nil {
			log.Fatal("Error reading file: " + err.Error())
		}

		// Upload the file
		_, err = serviceClient.UploadFile(ctx, containerName, blobName, file, nil)

		if err != nil {
			log.Fatal("Error creating blob: " + err.Error())
		}

		log.Printf("File '%s' uploaded", blobName)
	},
}

func init() {
	createCmd.AddCommand(blobCmd)

	// Flags
	blobCmd.Flags().StringP("name", "n", "", "Name of the blob")
	blobCmd.MarkFlagRequired("name")
	blobCmd.Flags().StringP("path", "p", "", "Local path of the blob")
	blobCmd.MarkFlagRequired("path")
	blobCmd.Flags().StringP("container", "c", "", "Container name")
	blobCmd.MarkFlagRequired("container")
}
