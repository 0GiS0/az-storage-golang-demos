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

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a blob",
	Long:  `You can download a blob from your storage account. You need to specify the container name, the blob name and the path where you want to save the file.`,
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

		file, err := os.Create(path)

		if err != nil {
			log.Fatal("Error reading file: " + err.Error())
		}

		// Download the blob
		_, err = serviceClient.DownloadFile(ctx, containerName, blobName, file, nil)

	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Flags
	downloadCmd.Flags().StringP("name", "n", "", "Name of the blob")
	downloadCmd.MarkFlagRequired("name")
	downloadCmd.Flags().StringP("path", "p", "", "Local path of the blob")
	downloadCmd.MarkFlagRequired("path")
	downloadCmd.Flags().StringP("container", "c", "", "Container name")
	downloadCmd.MarkFlagRequired("container")
}
