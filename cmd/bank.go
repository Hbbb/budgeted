package cmd

import (
	"bdgt/pkg/banks"
	"fmt"

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
	Run: func(cmd *cobra.Command, args []string) {
		bankName := args[0]
		publicToken := args[1]

		clientID := viper.GetString("plaid-client-id")
		publicKey := viper.GetString("plaid-public-key")
		secret := viper.GetString("plaid-secret")
		bankClient, err := banks.NewBankClient(clientID, publicKey, secret)
		if err != nil {
			panic(err)
		}

		accessToken, err := bankClient.ExchangePublicToken(publicToken)
		if err != nil {
			panic(err)
		}

		fmt.Println(accessToken)
		banks.Add(bankName, accessToken)
	},
}

var bankRemoveCmd = &cobra.Command{
	SilenceUsage: true,
	Use:          "remove [bankName]",
	Args:         cobra.ExactArgs(1),
	Short:        "removes bank from manifest",
	Run: func(cmd *cobra.Command, args []string) {
		bankName := args[0]

		err := banks.Remove(bankName)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(bankCmd)
	bankCmd.AddCommand(bankAddCmd, bankRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bankCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bankCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
