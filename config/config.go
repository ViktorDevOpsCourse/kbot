package config

import "os"

var (
	Version          = "v1.0.2"
	TelegramBotToken = os.Getenv("TELE_TOKEN")
)
