package cmd

import (
	"bdgt/pkg/banks"
	"fmt"
	"os"
	"text/tabwriter"
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

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "transaction id\t account id\t date\t name\t amount\t city")
		for _, transaction := range transactions {
			fmt.Fprintf(w, "%v\t %v\t %v\t %v\t %v\t %v", transaction.ID, transaction.AccountID, transaction.Date, transaction.Name, transaction.Amount, transaction.City)
			fmt.Fprintln(w, "")
		}
		w.Flush()

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
