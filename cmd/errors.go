package cmd

import "errors"

var (
	errMissingClientID      = errors.New("missing plaid client ID")
	errMissingPublicKey     = errors.New("missing plaid public key")
	errMissingSecret        = errors.New("missing plaid secret")
	errMissingPlaidEnv      = errors.New("missing plaid env")
	errMissingSpreadsheetID = errors.New("missing spreadsheet ID")
	errNotConfigured        = errors.New("not configured. run `config` command to configure bdgt")
)
