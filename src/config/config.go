package config

import (
	"os"
)

var (
	Port     = os.Getenv("PORT")
	BotToken = os.Getenv("TOKEN")
)
