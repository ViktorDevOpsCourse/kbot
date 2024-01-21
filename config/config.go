package config

import "os"

var (
	Version          = "Unknown version"
	TelegramBotToken = os.Getenv("TELE_TOKEN")

	// MetricsHost exporter host:port
	MetricsHost = os.Getenv("METRICS_HOST")
)
