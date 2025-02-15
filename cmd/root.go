package cmd

import (
	"bdgt/pkg/core"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:           "bdgt",
	Short:         "Bdgt is a tool to pull your bank transactions into a Google Spreadsheet",
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute executes a command
func Execute() {
	godotenv.Load()

	configPath, err := core.ConfigPath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = errNotConfigured
		}
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
