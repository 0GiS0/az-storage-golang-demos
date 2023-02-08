/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azure-storage-golang-demos/internal"

	"context"
	"log"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// blobsCmd represents the blobs command
var blobsCmd = &cobra.Command{
	Use:   "blobs",
	Short: "List blobs in a container",
	Long:  `List blobs in a specific container.`,
	PreRun: func(cmd *cobra.Command, args []string) {

		internal.CheckCredentials()

	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// Get flags values
		containerName := cmd.Flag("container").Value.String()

		// Get blob service client
		serviceClient, err := internal.GetBlobServiceClient()

		if err != nil {
			log.Fatal("Error trying to get service client: " + err.Error())
		}

		// Get blobs in the container
		pager := serviceClient.NewListBlobsFlatPager(containerName, nil)

		// Define output table
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		table := table.New("Name", "Last modified")
		table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		// List blobs
		for pager.More() {

			resp, err := pager.NextPage(ctx)
			if err != nil {
				log.Fatal("Error listing containers: " + err.Error())
				return
			}
			for _, _blob := range resp.Segment.BlobItems {
				table.AddRow(*_blob.Name, _blob.Properties.LastModified)
			}
		}

		// Print output table
		table.Print()

	},
}

func init() {
	listCmd.AddCommand(blobsCmd)

	//Flags
	blobsCmd.Flags().StringP("container", "c", "", "Name of the container")
}
