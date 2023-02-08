/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"azure-storage-golang-demos/internal"

	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// containersCmd represents the containers command
var containersCmd = &cobra.Command{
	Use:   "containers",
	Short: "List containers in your storage account",
	Long:  `You can see all containers in your storage account and the type of access configure.`,
	PreRun: func(cmd *cobra.Command, args []string) {

		internal.CheckCredentials()

	},
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		// Get blob service client
		serviceClient, err := internal.GetBlobServiceClient()

		if err != nil {
			log.Fatal("Error trying to get service client: " + err.Error())
		}

		// List containers
		pager := serviceClient.NewListContainersPager(&azblob.ListContainersOptions{
			Include: azblob.ListContainersInclude{Metadata: true, Deleted: true},
		})

		// Define output table
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		table := table.New("Name", "Last modified")
		table.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

		for pager.More() {

			resp, err := pager.NextPage(ctx)
			if err != nil {
				log.Fatal("Error listing containers: " + err.Error())
				return
			}
			for _, _container := range resp.ContainerItems {
				table.AddRow(*_container.Name, _container.Properties.LastModified)
			}
		}

		table.Print()

	},
}

func init() {
	listCmd.AddCommand(containersCmd)
}
