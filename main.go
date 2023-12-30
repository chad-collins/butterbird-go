package main

import (
	"github.com/chad-collins/butterbird-go/internal/bot"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/logger"
)

func main() {
	// Read the configuration
	err := config.ReadConfig()
	if err != nil {
		logger.Fatal(err, "Failed to read configuration")
	}

	// Initialize the bot
	b := bot.NewBot()

	// Start the bot
	err = b.Start()
	if err != nil {
		logger.Fatal(err, "Failed to start bot")
	}

	// Keep the program running until manually terminated
	<-make(chan struct{})
}
