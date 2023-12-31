package config

import (
	"encoding/json"
	"os"

	"github.com/chad-collins/butterbird-go/internal/logger"
)

var (
	DiscordToken   string
	OpenAiToken    string
	BotPrefix      string
	BotName        string
	GPTClient      string
	GPTClientToken string
	GPTClients     map[string]GPTClientConfig
)

type ConfigStruct struct {
	DiscordToken string                     `json:"DiscordToken"`
	BotPrefix    string                     `json:"BotPrefix"`
	BotName      string                     `json:"BotName"`
	GPTClient    string                     `json:"GPTClient"`
	GPTClients   map[string]GPTClientConfig `json:"GPTClients"`
}

type GPTClientConfig struct {
	Token string `json:"Token"`
}

func ReadConfig() {
	logger.Info("Reading config file...")

	file, err := os.ReadFile("./config.json")
	if err != nil {
		logger.Fatal(err, "Could not read config file")
	}

	config := &ConfigStruct{} // Use ConfigStruct here
	err = json.Unmarshal(file, config)
	if err != nil {
		logger.Fatal(err, "Could not unmarshal config data")
	}

	DiscordToken = config.DiscordToken
	BotPrefix = config.BotPrefix
	BotName = config.BotName
	GPTClient = config.GPTClient
	GPTClients = config.GPTClients
	GPTClientToken = config.GPTClients[config.GPTClient].Token //not using right now. Hardcoding token reference into client

	logger.Info("Config Loaded")
}
