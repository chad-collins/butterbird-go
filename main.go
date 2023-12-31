package main

import (
	"github.com/chad-collins/butterbird-go/internal/bot"
)

func main() {

	b := bot.NewBot()
	b.Start()

	<-make(chan struct{}) // Block forever
}
