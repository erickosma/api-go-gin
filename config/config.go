package config

import (
	"os"
)

type Config struct {
	GoogleSheetsCredentials string
	WhatsAppToken           string
	SpreadsheetID           string
	ServerPort              string
}

func NewConfig() *Config {
	return &Config{
		GoogleSheetsCredentials: getEnv("GOOGLE_SHEETS_CREDENTIALS", ""),
		WhatsAppToken:           getEnv("WHATSAPP_TOKEN", ""),
		SpreadsheetID:           getEnv("SPREADSHEET_ID", ""),
		ServerPort:              getEnv("PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
