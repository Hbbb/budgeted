package cmd

import (
	"bdgt/pkg/banks"
	"bdgt/pkg/spreadsheet"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var importCmd = &cobra.Command{
	Use:          "import",
	Short:        "writes transactions directly to a Google Sheet",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		clientID := viper.GetString("plaid_client_id")
		publicKey := viper.GetString("plaid_public_key")
		secret := viper.GetString("plaid_secret")
		spreadsheetID := viper.GetString("spreadsheet_id")

		if clientID == "" {
			return errMissingClientID
		}

		if publicKey == "" {
			return errMissingPublicKey
		}

		if secret == "" {
			return errMissingSecret
		}

		if spreadsheetID == "" {
			return errMissingSpreadsheetID
		}

		bankClient, err := banks.NewBankClient(clientID, publicKey, secret)
		if err != nil {
			return err
		}

		transactions, err := bankClient.FetchTransactions(startDate, endDate)
		if err != nil {
			return err
		}

		var writeData [][]interface{}
		for _, tr := range transactions {
			row := []interface{}{tr.ID, tr.AccountID, tr.Date, tr.Name, tr.Amount, tr.City}
			writeData = append(writeData, row)
		}

		writer := spreadsheet.Writer{
			SpreadsheetID: spreadsheetID,
		}

		err = writer.Write(writeData)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	isoFormat := "2006-01-02"
	start := time.Now()
	end := start.Add(24 * time.Hour)

	importCmd.Flags().StringVar(&startDate, "start", start.Format(isoFormat), "start date formatted YYYY-MM-DD")
	importCmd.Flags().StringVar(&endDate, "end", end.Format(isoFormat), "end date formatted YYYY-MM-DD")
}
