package gapi

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func ReadFromSheet(srv *sheets.Service, spreadsheetId string, readRange string) (*sheets.ValueRange, error) {
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("[ReadFromSheet] unable to retrieve data from sheet: %w", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("[ReadFromSheet] no data found in the table")
	}

	return resp, nil
}
