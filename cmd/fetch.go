package cmd

import (
	"bdgt/pkg/banks"
	"fmt"
	"time"

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

		fmt.Println("date name amount")
		for _, transaction := range transactions {
			fmt.Printf("%v %v - %v\n", transaction.Date, transaction.Name, transaction.Amount)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	isoFormat := "2006-01-02"
	start := time.Now()
	end := start.Add(24 * time.Hour)

	fetchCmd.Flags().StringVarP(&startDate, "start-date", "s", start.Format(isoFormat), "start date formatted YYYY-MM-DD")
	fetchCmd.Flags().StringVarP(&endDate, "end-date", "e", end.Format(isoFormat), "end date formatted YYYY-MM-DD")
	fetchCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file path")
}
