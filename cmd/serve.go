package cmd

import (
	"bdgt/pkg/web"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var publicKey string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a local webserver that hosts the Plaid authentication UI",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		plaidEnv := viper.GetString("plaid-env")

		if len(publicKey) < 1 {
			return fmt.Errorf("missing plaid public key")
		}

		fmt.Println(`Open http://localhost:80 and follow the steps to get Plaid bank credentials.
When you're finished, kill the process running this server.`)

		return web.Serve(plaidEnv, publicKey)
	},
}

func init() {
	viper.SetDefault("plaid_env", "development")

	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	serveCmd.Flags().StringVarP(&publicKey, "key", "k", viper.GetString("plaid-public-key"), "plaid public key")
}
