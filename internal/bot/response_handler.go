package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
)

// GetMessageHandler returns a message handler function based on the message.
func GetMessageHandler(m *discordgo.MessageCreate) func() (string, bool, string) {
	if m.Author.Bot {
		return func() (string, bool, string) {
			return "", false, ""
		}
	}

	if m.Author.Username == "kroby1" && !strings.HasPrefix(m.Content, config.BotPrefix) && !strings.HasPrefix(m.Content, "hey "+config.BotPrefix) {
		return func() (string, bool, string) {
			return gpt.KrobyPrompt(), true, m.Content
		}
	}

	heyPrefix := "hey " + config.BotPrefix
	if strings.HasPrefix(m.Content, heyPrefix) {
		messageContent := trimPrefixFromMessage(m.Content, heyPrefix)
		return func() (string, bool, string) {
			return gpt.GeneralPrompt(), true, messageContent
		}
	}

	if strings.HasPrefix(m.Content, config.BotPrefix) {
		messageContent := trimPrefixFromMessage(m.Content, config.BotPrefix)
		return func() (string, bool, string) {
			return "", false, messageContent
		}
	}

	return func() (string, bool, string) {
		return "", false, ""
	}
}

// trimPrefixFromMessage trims the configured prefix from the message content.
func trimPrefixFromMessage(content, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(content, prefix))
}
