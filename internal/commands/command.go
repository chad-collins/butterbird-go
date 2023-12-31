package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/gpt"
)

// Command defines the interface for bot commands.
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate) error
	Init(gptClient gpt.GPTClient)
	Trigger() string
}
