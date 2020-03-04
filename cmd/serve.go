package cmd

import (
	"bdgt/pkg/web"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:          "serve",
	Short:        "start a local webserver that hosts the Plaid authentication UI",
	Long:         ``,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		plaidEnv := viper.GetString("plaid_env")
		publicKey := viper.GetString("plaid_public_key")

		fmt.Println(`Open http://localhost:80 and follow the steps to get Plaid bank credentials.
When you're finished, kill the process running this server.`)

		return web.Serve(plaidEnv, publicKey)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
