/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/spf13/cobra"

	"azure-storage-golang-demos/internal"
)

// containerCmd represents the container command
var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "It creates a container",
	Long:  `It creates a container in a storage account.`,
	PreRun: func(cmd *cobra.Command, args []string) {

		internal.CheckCredentials()

	},
	Run: func(cmd *cobra.Command, args []string) {

		var containerName string

		containerName, _ = cmd.Flags().GetString("name")
		isPublic, _ := cmd.Flags().GetBool("public")

		ctx := context.Background()

		// Get blob service client
		serviceClient, err := internal.GetBlobServiceClient()

		if err != nil {
			log.Fatal("Error trying to get service client: " + err.Error())
		}

		// Create a container
		containerClient := serviceClient.ServiceClient().NewContainerClient(containerName)

		if isPublic {

			log.Println("Creating public container")

			access := container.PublicAccessTypeContainer
			// access := container.PublicAccessTypeBlob

			createContainerOptions := container.CreateOptions{
				Access: &access,
			}
			_, err = containerClient.Create(ctx, &createContainerOptions)

		} else {
			_, err = containerClient.Create(ctx, nil)
		}

		if err != nil {
			log.Fatal("Error creating container: " + err.Error())
		}

		log.Printf("Container %s created", containerName)

	},
}

func init() {
	createCmd.AddCommand(containerCmd)

	// Here you will define your flags and configuration settings.
	containerCmd.Flags().StringP("name", "n", "", "Name of the container")
	containerCmd.MarkFlagRequired("name")
	containerCmd.Flags().BoolP("public", "p", false, "Make the container public")
}
