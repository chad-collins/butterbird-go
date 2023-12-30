package bot

import (
	"errors"
	"strings"

	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/logger"
	"github.com/chad-collins/butterbird-go/internal/utils/openAiUtils"

	"github.com/bwmarrin/discordgo"
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
		logger.Fatal(err, "Failed to create Discord session")
	}
	logger.Info("Discord session created successfully")

	openAiClient := openai.NewClient(config.OpenAiToken)
	if openAiClient == nil {
		logger.Fatal(errors.New("OpenAI client initialization failed"), "Failed to create OpenAI session")
	}
	logger.Info("OpenAI client initialized successfully")

	bot := &Bot{
		Session:      session,
		OpenAiClient: openAiClient,
	}

	user, err := bot.Session.User("@me")
	if err != nil {
		logger.Fatal(err, "Failed to access bot user details")
	}
	bot.ID = user.ID
	logger.Info("Bot user details retrieved successfully")

	return bot
}

// Start begins the bot's operation.
func (b *Bot) Start() error {
	b.Session.AddHandler(b.messageCreate)
	b.Session.Identify.Intents = discordgo.IntentsGuildMessages

	err := b.Session.Open()
	if err != nil {
		return err
	}
	logger.Info("Websocket connection to Discord opened successfully")
	logger.Info(config.BotName + " is online!")
	return nil
}

func (b *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	logger.Info("Message received")

	// Ignore messages sent by bots, including itself
	if m.Author.Bot {
		logger.Info("Message from a bot, ignored.")
		return
	}

	messageContent := m.Content
	var prompt string

	// Check if the message is from "kroby1" and does not start with specific prefixes
	if m.Author.Username == "kroby1" && (!strings.HasPrefix(m.Content, "hey "+config.BotPrefix) || !strings.HasPrefix(m.Content, config.BotPrefix)) {
		logger.Info("Special processing for kroby1 without specific prefixes")
		prompt = openAiUtils.KrobyPrompt()
	} else if strings.HasPrefix(m.Content, "hey "+config.BotPrefix) || strings.HasPrefix(m.Content, config.BotPrefix) {
		logger.Info("Message starts with the prefix")
		// Remove the prefix for processing
		messageContent = strings.TrimSpace(strings.TrimPrefix(m.Content, "hey "+config.BotPrefix))
		prompt = openAiUtils.GeneralPrompt()
	} else {
		return
	}

	// Check if the message content is not empty
	if len(messageContent) == 0 {
		logger.Info("No content in message to process")
		return
	}

	// Process the message content with OpenAI
	res, err := openAiUtils.BuildPrompt(b.OpenAiClient, messageContent, prompt, "")
	if err != nil {
		logger.Warn(err, "Failed to transform message")
		return
	}
	logger.Info("Message transformed successfully")

	// Send the transformed message
	if _, err := s.ChannelMessageSend(m.ChannelID, res); err != nil {
		logger.Warn(err, "Failed to send message")
		return
	}
	logger.Info("Transformed message sent successfully")
}
