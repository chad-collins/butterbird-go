package bot

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/logger"
	"github.com/chad-collins/butterbird-go/internal/utils/openAiUtils"
	"github.com/sashabaranov/go-openai"
)

// Bot represents the state of the Discord bot.
type Bot struct {
	ID           string
	Session      *discordgo.Session
	OpenAiClient *openai.Client
}

// NewBot initializes a new Bot instance.
func NewBot() *Bot {
	session, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		logger.Fatal(err, "Creating Discord session")
	}
	logger.Info("Discord session created")

	openAiClient := openai.NewClient(config.OpenAiToken)
	if openAiClient == nil {
		logger.Fatal(errors.New("initialization failure"), "Creating OpenAI client")
	}
	logger.Info("OpenAI client initialized")

	bot := &Bot{
		Session:      session,
		OpenAiClient: openAiClient,
	}

	user, err := bot.Session.User("@me")
	if err != nil {
		logger.Fatal(err, "Accessing bot user details")
	}
	bot.ID = user.ID
	logger.Info("Bot user details retrieved")

	return bot
}

// Start begins the bot's operation.
func (b *Bot) Start() {
	b.Session.AddHandler(b.messageCreate)
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages

	if err := b.Session.Open(); err != nil {
		logger.Fatal(err, "Opening WebSocket connection")
	}
	logger.Info(config.BotName + " is online!")
}

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return // Ignore bot messages
	}

	messageContent := m.Content
	var prompt string

	if m.Author.Username == "kroby1" && !strings.HasPrefix(m.Content, "hey "+config.BotPrefix) && !strings.HasPrefix(m.Content, config.BotPrefix) {
		prompt = openAiUtils.KrobyPrompt()
	} else if strings.HasPrefix(m.Content, "hey "+config.BotPrefix) || strings.HasPrefix(m.Content, config.BotPrefix) {
		messageContent = strings.TrimSpace(strings.TrimPrefix(m.Content, "hey "+config.BotPrefix))
		prompt = openAiUtils.GeneralPrompt()
	} else {
		return
	}

	if len(messageContent) == 0 {
		return // No content to process
	}

	res, err := openAiUtils.BuildPrompt(b.OpenAiClient, messageContent, prompt, "")
	if err != nil {
		logger.Warn(err, "Transforming message")
		return
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, res); err != nil {
		logger.Warn(err, "Sending transformed message")
		return
	}
}
