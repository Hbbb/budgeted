package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bdgt",
	Short: "Bdgt is a tool to pull your bank transactions into a Google Spreadsheet",
}

// Execute executes a command
func Execute() {
	godotenv.Load()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
