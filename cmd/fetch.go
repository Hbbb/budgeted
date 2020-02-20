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

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch transactions",
	Run: func(cmd *cobra.Command, args []string) {
		clientID := viper.GetString("plaid-client-id")
		publicKey := viper.GetString("plaid-public-key")
		secret := viper.GetString("plaid-secret")

		bankClient, err := banks.NewBankClient(clientID, publicKey, secret)
		if err != nil {
			panic(err)
		}

		transactions, err := bankClient.FetchTransactions(startDate, endDate)
		if err != nil {
			panic(err)
		}

		for _, transaction := range transactions {
			fmt.Printf("%v %v - %v\n", transaction.Date, transaction.Name, transaction.Amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	isoFormat := "2006-01-02"
	start := time.Now()
	end := start.Add(24 * time.Hour)

	fetchCmd.Flags().StringVarP(&startDate, "start-date", "s", start.Format(isoFormat), "start date formatted YYYY-MM-DD")
	fetchCmd.Flags().StringVarP(&endDate, "end-date", "e", end.Format(isoFormat), "end date formatted YYYY-MM-DD")
}
