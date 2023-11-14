package config

import "os"

var (
	Version          = "v1.0.1"
	TelegramBotToken = os.Getenv("TELE_TOKEN")
)
