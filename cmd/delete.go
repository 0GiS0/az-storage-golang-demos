/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azure-storage-golang-demos/internal"
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a blob from a container",
	Long:  `You can delete blobs in a container.`,
	PreRun: func(cmd *cobra.Command, args []string) {

		internal.CheckCredentials()

	},
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		// Get flags values
		containerName := cmd.Flag("container").Value.String()
		blobName := cmd.Flag("name").Value.String()

		// Get blob service client
		serviceClient, err := internal.GetBlobServiceClient()

		if err != nil {
			log.Fatal("Error trying to get service client: " + err.Error())
		}
		serviceClient.DeleteBlob(ctx, containerName, blobName, nil)

		fmt.Printf("Blob %s deleted from container %s successfully", blobName, containerName)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.
	deleteCmd.Flags().StringP("name", "n", "", "Name of the blob")
	deleteCmd.MarkFlagRequired("name")
	deleteCmd.Flags().StringP("container", "c", "", "Container name")
	deleteCmd.MarkFlagRequired("container")
}
