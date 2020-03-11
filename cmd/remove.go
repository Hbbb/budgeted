package cmd

import (
	"bdgt/pkg/banks"

	"github.com/spf13/cobra"
)

var bankRemoveCmd = &cobra.Command{
	Use:   "remove [bankName]",
	Args:  cobra.ExactArgs(1),
	Short: "removes bank from manifest",
	RunE: func(cmd *cobra.Command, args []string) error {
		bankName := args[0]
		return banks.Remove(bankName)
	},
}

func init() {
	rootCmd.AddCommand(bankRemoveCmd)
}
