package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
)

type KrobyCommand struct {
	GPTClient gpt.GPTClient
}

func (cmd *KrobyCommand) Init(gptClient gpt.GPTClient) {
	cmd.GPTClient = gptClient
}

func (cmd KrobyCommand) Trigger() string {
	return ""
}

func (cmd KrobyCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) error {
	if m.Author.Username == "kroby1" &&
		!strings.HasPrefix(m.Content, "hey "+config.BotPrefix) &&
		!strings.HasPrefix(m.Content, config.BotPrefix) {

		messageContent := strings.TrimSpace(m.Content)
		response, err := cmd.GPTClient.BuildPrompt(messageContent, krobyPrompt(), "")
		if err != nil {
			// Send a user-friendly error message back to the channel
			s.ChannelMessageSend(m.ChannelID, "Sorry, I encountered an error processing your request.")
			return err
		}

		_, err = s.ChannelMessageSend(m.ChannelID, response)
		return err
	}
	return nil
}

func krobyPrompt() string {
	return "Please assist in rewriting the following message from Kroby. The original author has difficulty typing, often uses incorrect words, and doesn't use punctuation, which can lead to confusion in their meaning. Your task is to correct all grammar, rectify any incorrect word usage, and add proper punctuation. Also, adjust the sentence structure to enhance readability and ensure the intended meaning is clear:\n"
}
