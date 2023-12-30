package config

import (
	"encoding/json"
	"os"

	"github.com/chad-collins/butterbird-go/internal/logger"
)

var (
	DiscordToken string
	OpenAiToken  string
	BotPrefix    string
	BotName      string
)

type configStruct struct {
	DiscordToken string `json:"DiscordToken"`
	OpenAiToken  string `json:"OpenAiToken"`
	BotPrefix    string `json:"BotPrefix"`
	BotName      string `json:"BotName"`
}

func ReadConfig() {
	logger.Info("Reading config file...")

	file, err := os.ReadFile("./config.json")
	if err != nil {
		logger.Fatal(err, "Could not read config file")
	}

	config := &configStruct{}
	err = json.Unmarshal(file, config)
	if err != nil {
		logger.Fatal(err, "Could not unmarshal config data")
	}

	DiscordToken = config.DiscordToken
	OpenAiToken = config.OpenAiToken
	BotPrefix = config.BotPrefix
	BotName = config.BotName

	logger.Info("Config Loaded")
}
