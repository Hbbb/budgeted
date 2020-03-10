package cmd

import (
	"bdgt/pkg/core"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure Plaid API keys and Google Sheets Spreadsheet ID",
	Long: `
Find your Plaid API keys and secrets here: https://dashboard.plaid.com/overview/development
Find your spreadsheet ID in the URL of your Google Sheet e.g. https://docs.google.com/spreadsheets/d/<SPREADSHEET_ID>/
	`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var clientID string
		var publicKey string
		var secret string
		var spreadsheetID string

		fmt.Print("Plaid Client ID: ")
		_, err := fmt.Scan(&clientID)
		if err != nil {
			return err
		}

		fmt.Print("Plaid Public Key: ")
		_, err = fmt.Scan(&publicKey)
		if err != nil {
			return err
		}

		fmt.Print("Plaid Secret: ")
		_, err = fmt.Scan(&secret)
		if err != nil {
			return err
		}

		fmt.Print("Spreadsheet ID: ")
		_, err = fmt.Scan(&spreadsheetID)
		if err != nil {
			return err
		}

		viper.Set("plaid_client_id", clientID)
		viper.Set("plaid_public_key", publicKey)
		viper.Set("plaid_secret", secret)
		viper.Set("plaid_env", "development")
		viper.Set("spreadsheet_id", spreadsheetID)

		err = viper.WriteConfig()
		if err != nil {
			return err
		}

		configDir, err := core.ConfigPath()
		if err != nil {
			return err
		}

		_, err = os.Stat(configDir + "/banks.json")
		if os.IsNotExist(err) {
			err = ioutil.WriteFile(configDir+"/banks.json", []byte{}, 0644)
			if err != nil {
				return err
			}
		}

		fmt.Println("Writing banks.json")
		fmt.Println("Configuration saved to ~/.budgeted/config.yaml")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
