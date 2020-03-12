package sheets

import "errors"

var (
	errSheetsAPINotEnabled = errors.New("sheets: Google Sheets API not enabled. enable the Google Sheets API here: https://developers.google.com/sheets/api/quickstart/go")
	errSheetsAuthFailure   = errors.New("sheets: authorization failed. try again")
	errWriteFailed         = errors.New("sheets: writing to spreadhseet failed")
)
