package cmd

import (
	"bdgt/pkg/banks"
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startDate string
var endDate string
var outputFile string

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch transactions",
	RunE: func(cmd *cobra.Command, args []string) error {
		clientID := viper.GetString("plaid_client_id")
		publicKey := viper.GetString("plaid_public_key")
		secret := viper.GetString("plaid_secret")

		if clientID == "" {
			return errMissingClientID
		}

		if publicKey == "" {
			return errMissingPublicKey
		}

		if secret == "" {
			return errMissingSecret
		}

		// NOTE: This always uses Plaid development env, which I guess is fine for now
		bankClient, err := banks.NewBankClient(clientID, publicKey, secret)
		if err != nil {
			return err
		}

		transactions, err := bankClient.FetchTransactions(startDate, endDate)
		if err != nil {
			return err
		}

		if len(outputFile) > 0 {
			tbytes, err := json.Marshal(transactions)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(outputFile, tbytes, 0644)
			if err != nil {
				return err
			}

			return nil
		}

		w := tablewriter.NewWriter(os.Stdout)
		w.SetHeader([]string{"transaction id", "account id", "date", "name", "amount", "city"})
		for _, transaction := range transactions {
			w.Append([]string{transaction.ID, transaction.AccountID, transaction.Date, transaction.Name, transaction.Amount, transaction.City})
		}
		w.Render()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	isoFormat := "2006-01-02"
	start := time.Now()
	end := start.Add(24 * time.Hour)

	fetchCmd.Flags().StringVar(&startDate, "start", start.Format(isoFormat), "start date formatted YYYY-MM-DD")
	fetchCmd.Flags().StringVar(&endDate, "end", end.Format(isoFormat), "end date formatted YYYY-MM-DD")
	fetchCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file path")
}
