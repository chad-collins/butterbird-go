package main

import (
	"github.com/chad-collins/butterbird-go/internal/bot"
	"github.com/chad-collins/butterbird-go/internal/config"
)

func main() {
	config.ReadConfig()
	b := bot.NewBot()
	b.Start()

	<-make(chan struct{}) // Block forever
}
