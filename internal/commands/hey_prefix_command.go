package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
)

type HeyPrefixCommand struct {
	GPTClient gpt.GPTClient
}

func (cmd *HeyPrefixCommand) Init(gptClient gpt.GPTClient) {
	cmd.GPTClient = gptClient
}

func (cmd HeyPrefixCommand) Trigger() string {
	return "hey " + config.BotPrefix
}

func (cmd HeyPrefixCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) error {
	messageContent := strings.TrimSpace(strings.TrimPrefix(m.Content, "hey "+config.BotPrefix))
	response, err := cmd.GPTClient.BuildPrompt(messageContent, generalPrompt(), "")
	if err != nil {
		// Send a user-friendly error message back to the channel
		s.ChannelMessageSend(m.ChannelID, "Sorry, I encountered an error processing your request.")
		return err
	}

	_, err = s.ChannelMessageSend(m.ChannelID, response)
	return err
}

func generalPrompt() string {
	return "Please respond like Farva from Super Troopers:\n"
}
