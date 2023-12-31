package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/chad-collins/butterbird-go/internal/commands"
	"github.com/chad-collins/butterbird-go/internal/config"
	"github.com/chad-collins/butterbird-go/internal/gpt"
	"github.com/chad-collins/butterbird-go/internal/logger"
)

// Bot represents the state of the Discord bot.
type Bot struct {
	ID        string
	Session   *discordgo.Session
	GPTClient gpt.GPTClient // Use the GPTClient interface
	Commands  map[string]commands.Command
}

var clientFactories = map[string]func(token string) gpt.GPTClient{
	"openai": func(token string) gpt.GPTClient {
		return new(gpt.OpenAIGPTClient).CreateClient(token)
	},
}

// LoadCommands initializes and maps the commands for the bot.
func (b *Bot) LoadCommands() {
	// List of command instances
	commandList := []commands.Command{
		&commands.HelloCommand{},
		&commands.HeyPrefixCommand{},
		&commands.KrobyCommand{},
		// ... add other commands here ...
	}

	for _, cmd := range commandList {
		cmd.Init(b.GPTClient)           // Initialize the command
		b.Commands[cmd.Trigger()] = cmd // Load the command using its trigger
	}
}

// NewBot initializes a new Bot instance.
func NewBot() *Bot {
	config.ReadConfig() // Ensure configuration is loaded

	session, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		logger.Fatal(err, "Creating Discord session")
	}
	logger.Info("Discord session created")

	// Set the client type and get the corresponding token
	clientType := config.GPTClient
	gptToken := config.GPTClients[clientType].Token

	factory, exists := clientFactories[clientType]
	if !exists {
		logger.Fatal(nil, "Unknown GPT client type: "+clientType)
	}
	// Initialize the GPT client with the token
	gptClient := factory(gptToken)

	bot := &Bot{
		Session:   session,
		GPTClient: gptClient,
		Commands:  make(map[string]commands.Command),
	}

	user, err := session.User("@me")
	if err != nil {
		logger.Fatal(err, "Accessing bot user details")
	}
	bot.ID = user.ID
	logger.Info("Bot user details retrieved")

	bot.LoadCommands()

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
	// Ignore messages from bots to prevent potential loops or spam
	if m.Author.Bot {
		return
	}

	// Iterate over registered commands
	for trigger, cmd := range b.Commands {
		// Check if the message starts with the command trigger
		if strings.HasPrefix(m.Content, trigger) {
			// Execute the command and handle any errors
			err := cmd.Execute(s, m)
			if err != nil {
				logger.Warn(err, "Executing command")
			}
			return // Stop processing after the first matching command
		}
	}
}
