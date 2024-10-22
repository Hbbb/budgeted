package sheets

import gsheets "google.golang.org/api/sheets/v4"

const writeRange = "!A2:D"

// Writer writes
type Writer struct {
	SpreadsheetID string
}

// TODO: Handle errors
func (w *Writer) Write(data [][]interface{}) error {
	srv, err := newSheetsService()
	if err != nil {
		return err
	}

	var vr gsheets.ValueRange
	vr.Range = writeRange

	for _, values := range data {
		vr.Values = append(vr.Values, values)
	}

	_, err = srv.Spreadsheets.Values.
		Append(w.SpreadsheetID, writeRange, &vr).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return errWriteFailed
	}

	return nil
}
