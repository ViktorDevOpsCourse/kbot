package config

import "os"

var (
	Version          = "Unknown version"
	TelegramBotToken = os.Getenv("TELE_TOKEN")
)
