package cmd

import (
	"bdgt/pkg/banks"
	"bdgt/pkg/spreadsheet"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "writes transactions directly to a Google Sheet",
	RunE: func(cmd *cobra.Command, args []string) error {
		clientID := viper.GetString("plaid-client-id")
		publicKey := viper.GetString("plaid-public-key")
		secret := viper.GetString("plaid-secret")

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

	importCmd.Flags().StringVarP(&spreadsheetID,
		"spreadsheet-id", "s",
		viper.GetString("spreadsheet-id"),
		"the ID of spreadsheet to write transaction data to; defaults to SPREADSHEET_ID env var")
}