package spreadsheet

import "google.golang.org/api/sheets/v4"

const writeRange = "!A:D"

// Writer writes
type Writer struct {
	SpreadsheetID string
}

func (w *Writer) Write(data [][]interface{}) error {
	srv, err := newSheetsService()

	var vr sheets.ValueRange

	// Write headers
	vr.Values = append(vr.Values, []interface{}{"transaction id", "account id", "date", "name", "amount", "city"})

	for _, values := range data {
		vr.Values = append(vr.Values, values)
	}

	_, err = srv.Spreadsheets.Values.Update(w.SpreadsheetID, writeRange, &vr).Do()
	if err != nil {
		return err
	}

	return nil
}
