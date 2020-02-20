package cmd

import (
	"github.com/spf13/cobra"
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

	},
}

var bankRemoveCmd = &cobra.Command{
	SilenceUsage: true,
	Use:          "remove [bankName]",
	Args:         cobra.ExactArgs(2),
	Short:        "removes bank from manifest",
	Run: func(cmd *cobra.Command, args []string) {

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
