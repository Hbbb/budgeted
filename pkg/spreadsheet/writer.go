package spreadsheet

import "google.golang.org/api/sheets/v4"

const writeRange = "!A2:D"

// Writer writes
type Writer struct {
	SpreadsheetID string
}

func (w *Writer) Write(data [][]interface{}) error {
	srv, err := newSheetsService()

	var vr sheets.ValueRange
	vr.Range = writeRange

	// Write headers
	vr.Values = append(vr.Values, []interface{}{"Transaction ID", "Account ID", "Date", "Name", "Amount", "City"})

	for _, values := range data {
		vr.Values = append(vr.Values, values)
	}

	_, err = srv.Spreadsheets.Values.
		Append(w.SpreadsheetID, writeRange, &vr).
		ValueInputOption("USER_ENTERED").
		Do()

	if err != nil {
		return err
	}

	return nil
}
