package bot

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
	"github.com/chad-collins/butterbird-go/internal/logger"
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
	b.Session.AddHandler(b.OnMessageReceived)
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages

	if err := b.Session.Open(); err != nil {
		logger.Fatal(err, "Opening WebSocket connection")
	}
	logger.Info(config.BotName + " is online!")
}

// OnMessageReceived is called when a message is received.
func (b *Bot) OnMessageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	handler := GetMessageHandler(m)

	prompt, shouldHandle, messageContent := handler()

	if !shouldHandle {
		return // Ignore non-valid messages
	}

	res, err := gpt.BuildPrompt(b.OpenAiClient, messageContent, prompt, "")
	if err != nil {
		logger.Warn(err, "Transforming message")
		return
	}

	if _, err := s.ChannelMessageSend(m.ChannelID, res); err != nil {
		logger.Warn(err, "Sending transformed message")
		return
	}
}
