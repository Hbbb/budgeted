package cmd

import (
	"bdgt/pkg/banks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var bankAddCmd = &cobra.Command{
	Use:   "add [bankName] [publicToken]",
	Args:  cobra.ExactArgs(2),
	Short: "adds bank to manifest for transaction fetching",
	RunE: func(cmd *cobra.Command, args []string) error {
		bankName := args[0]
		publicToken := args[1]

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

		bankClient, err := banks.NewBankClient(clientID, publicKey, secret)
		if err != nil {
			return err
		}

		accessToken, err := bankClient.ExchangePublicToken(publicToken)
		if err != nil {
			return err
		}

		return banks.Add(bankName, accessToken)
	},
}

func init() {
	rootCmd.AddCommand(bankAddCmd)
}
