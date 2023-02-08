package internal

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

func CheckCredentials() {

	// Get config values if there is a config file in $HOME
	storageAccountName := viper.GetString("storageAccountName")

	if storageAccountName == "" {
		// Set config file
		viper.AddConfigPath(".")
		viper.SetConfigFile("config.yaml")
		viper.ReadInConfig()
	}

	// try again
	storageAccountName = viper.GetString("storageAccountName")

	// if not, prompt for azure storage account name
	if storageAccountName == "" {

		// Prompt for azure storage account name
		prompt := promptui.Prompt{
			Label: "Type your storage account name",
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)

		viper.AddConfigPath(".")
		viper.SetDefault("storageAccountName", result)
		viper.WriteConfig()
	}

	// Print config values
	println("You are playing with " + storageAccountName)
}
