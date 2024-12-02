package service

import (
	"api-go-gin/internal/domain"
)

const (
	SpreadsheetID = "<YOUR_SPREADSHEET_ID>"
	SheetRange    = "Sheet1!A:D" // Ajuste para o range correto
	Credentials   = "credentials.json"
)

func SaveToGoogleSheets(data domain.User) error {
	//ctx := context.Background()
	//srv, err := sheets.NewService(ctx, option.WithCredentialsFile(Credentials))
	//if err != nil {
	//	log.Fatalf("Unable to retrieve Sheets client: %v", err)
	//	return err
	//}
	//
	//values := [][]interface{}{
	//	{data.Name, data.Email, data.WhatsApp},
	//}
	//
	//rb := &sheets.ValueRange{
	//	Values: values,
	//}
	//
	//_, err = srv.Spreadsheets.Values.Append(SpreadsheetID, SheetRange, rb).
	//	ValueInputOption("RAW").
	//	Do()
	//if err != nil {
	//	return fmt.Errorf("unable to write data to sheet: %v", err)
	//}

	return nil
}
