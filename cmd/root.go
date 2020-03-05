package cmd

import (
	"bdgt/pkg/core"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:          "bdgt",
	Short:        "Bdgt is a tool to pull your bank transactions into a Google Spreadsheet",
	SilenceUsage: true,
}

// Execute executes a command
func Execute() {
	godotenv.Load()

	configPath, err := core.ConfigPath()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = createConfigFile(configPath)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createConfigFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	if _, err := os.Stat(path + "/config.yaml"); os.IsNotExist(err) {
		ioutil.WriteFile(path+"/config.yaml", []byte{}, 0644)
	}

	return nil
}
