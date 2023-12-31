package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
)

type HelloCommand struct {
	// No GPTClient needed in this command.
}

// Init does nothing for HelloCommand but is necessary to satisfy the Command interface.
func (cmd *HelloCommand) Init(gptClient gpt.GPTClient) {
	// No initialization needed for HelloCommand.
}

// Trigger returns the command trigger string.
func (cmd *HelloCommand) Trigger() string {
	return config.BotPrefix + " hello"
}

// Execute runs the command's logic.
func (cmd *HelloCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) error {
	_, err := s.ChannelMessageSend(m.ChannelID, "Hello there!")
	return err
}
