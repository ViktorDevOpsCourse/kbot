package config

import "os"

var Version = "v1.0.0"

var (
	TelegramBotToken = os.Getenv("TELE_TOKEN")
)
