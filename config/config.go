package config

import (
	"os"
)

type Config struct {
	GoogleSheetsCredentials string
	WhatsAppToken           string
	SpreadsheetID           string
}

func NewConfig() *Config {
	return &Config{
		GoogleSheetsCredentials: os.Getenv("GOOGLE_SHEETS_CREDENTIALS"),
		WhatsAppToken:           os.Getenv("WHATSAPP_TOKEN"),
		SpreadsheetID:           os.Getenv("SPREADSHEET_ID"),
	}
}
