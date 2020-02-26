package cmd

import (
	"bdgt/pkg/banks"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var bankCmd = &cobra.Command{
	Use: "bank [add || remove]",
}

var bankAddCmd = &cobra.Command{
	SilenceUsage: true,
	Use:          "add [bankName] [publicToken]",
	Args:         cobra.ExactArgs(2),
	Short:        "adds bank to manifest for transaction fetching",
	RunE: func(cmd *cobra.Command, args []string) error {
		bankName := args[0]
		publicToken := args[1]

		clientID := viper.GetString("plaid-client-id")
		publicKey := viper.GetString("plaid-public-key")
		secret := viper.GetString("plaid-secret")
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

var bankRemoveCmd = &cobra.Command{
	SilenceUsage: true,
	Use:          "remove [bankName]",
	Args:         cobra.ExactArgs(1),
	Short:        "removes bank from manifest",
	RunE: func(cmd *cobra.Command, args []string) error {
		bankName := args[0]
		return banks.Remove(bankName)
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
	bankCmd.AddCommand(bankAddCmd, bankRemoveCmd)
}
