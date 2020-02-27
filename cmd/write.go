package cmd

import (
	"bdgt/pkg/banks"
	"bdgt/pkg/spreadsheet"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// flag variables
var file string
var spreadsheetID string

var writeCmd = &cobra.Command{
	Use:           "write",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "write the contents of a JSON file of transaction data to a Google Sheet",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(file) < 1 {
			return fmt.Errorf("must pass a valid filepath")
		}

		// TODO: Check for file existence
		f, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		var transactions []banks.Transaction
		err = json.Unmarshal(f, &transactions)
		if err != nil {
			return err
		}

		writer := spreadsheet.Writer{
			SpreadsheetID: spreadsheetID,
		}

		var writeData [][]interface{}
		for _, tr := range transactions {
			row := []interface{}{tr.ID, tr.AccountID, tr.Date, tr.Name, tr.Amount, tr.City}
			writeData = append(writeData, row)
		}

		err = writer.Write(writeData)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	writeCmd.Flags().StringVarP(&file,
		"file", "f", "",
		"filepath to transaction data; generate the file by writing the output of the fetch command to a file")

	writeCmd.Flags().StringVarP(&spreadsheetID,
		"spreadsheet-id", "s",
		viper.GetString("spreadsheet-id"),
		"the ID of spreadsheet to write transaction data to; defaults to SPREADSHEET_ID env var")

	// TODO: Accept input from pipe
	// info, _ := os.Stdin.Stat()

	// if info.Mode()&os.ModeCharDevice == 0 {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	var output []rune

	// 	for {
	// 		input, _, err := reader.ReadRune()
	// 		if err != nil && err == io.EOF {
	// 			break
	// 		}
	// 		output = append(output, input)
	// 	}

	// 	// At this point, output is the full content of the data
	// 	for _, ch := range output {
	// 		fmt.Printf("%c", ch)
	// 	}

	// 	return nil
	// }
}
