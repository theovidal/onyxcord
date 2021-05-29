package onyxcord

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
)

// Bot represents the Discord bot with its assets
type Bot struct {
	// Name of the bot
	Name string
	// The Discord client associated with the bot
	Client *discordgo.Session
	// A list of all the slash commands available on the bot
	Commands map[string]*Command
	// A list of all the components handlers (buttons...)
	Components map[string]Component
	// The mongodb database attached to the bot (if used)
	Database *mongo.Client
	// The Redis cache attached to the bot (if used)
	Cache *redis.Client
	// The configuration of the bot, as defined in the corresponding file
	Config *Config
	// The profile of the bot
	User *discordgo.User
}

// RegisterBot creates a new instance of the Discord Bot
func RegisterBot(name string) Bot {
	// Loading the configuration
	config, err := GetConfig()
	if err != nil {
		log.Panicf("‼ Error getting the configuration : %v", err)
	}

	// Creating the bot profile
	client, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		panic(err)
	}
	client.Debug = config.Dev.Debug
	client.StateEnabled = true

	// Creating the user profile
	user, err := client.User(config.Bot.ID)
	if err != nil {
		log.Panicf("‼ Error creating the user profile : %v", err)
	}

	bot := Bot{
		Name:     name,
		Client:   client,
		Commands: make(map[string]*Command),
		Config:   &config,
		User:     user,
	}
	return bot
}

// RegisterCommand is a shorthand to add a slash command into the bot
func (bot *Bot) RegisterCommand(name string, command *Command) {
	bot.Commands[name] = command
}

// RegisterComponent is a shorthand to add a component handler into the bot
func (bot *Bot) RegisterComponent(name string, component Component) {
	bot.Components[name] = component
}

// Run starts the bot and connects it to Discord
func (bot *Bot) Run(registerHandler bool) {
	fmt.Println(" ________  ________       ___    ___ ___    ___ \n|\\   __  \\|\\   ___  \\    |\\  \\  /  /|\\  \\  /  /|\n\\ \\  \\|\\  \\ \\  \\\\ \\  \\   \\ \\  \\/  / | \\  \\/  / /\n \\ \\  \\\\\\  \\ \\  \\\\ \\  \\   \\ \\    / / \\ \\    / / \n  \\ \\  \\\\\\  \\ \\  \\\\ \\  \\   \\/  /  /   /     \\/  \n   \\ \\_______\\ \\__\\\\ \\__\\__/  / /    /  /\\   \\  \n    \\|_______|\\|__| \\|__|\\___/ /    /__/ /\\ __\\ \n                        \\|___|/     |__|/ \\|__| ")

	if bot.Config.Database.Address != "" {
		log.Println("⏩ Opening database...")
		bot.Database = OpenDatabase(bot.Config)
	}
	if bot.Config.Cache.Address != "" {
		log.Println("⏩ Opening cache...")
		bot.Cache = OpenCache(bot.Config)
	}

	if registerHandler {
		bot.Client.AddHandler(func(_ *discordgo.Session, interaction *discordgo.InteractionCreate) {
			err := bot.HandleInteraction(interaction)
			if err != nil {
				log.Println(Red.Sprintf("‼ Error responding to an interaction: %s", err))
			}
		})
	}

	log.Println("⏩ Connecting bot...")
	// Open a websocket connection to Discord and begin listening.
	err := bot.Client.Open()
	if err != nil {
		log.Fatalf("‼ Error opening connection with bot %s: %s\n", bot.Name, err)
	}
	log.Println(Green.Sprintf("✅ %s#%s logged in!", bot.User.Username, bot.User.Discriminator))

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("💤 Closing down bot...")

	bot.Client.Close()

	if bot.Config.Database.Address != "" {
		bot.Database.Disconnect(context.Background())
	}
	log.Println("👋 Goodbye!")
}

// HandleInteraction executes an interaction (or slash command)
func (bot *Bot) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	switch interaction.Type {
	case discordgo.InteractionApplicationCommand: return bot.HandleCommand(interaction)
	case discordgo.InteractionMessageComponent: return bot.HandleComponent(interaction)
	}
	return nil
}

func (bot *Bot) HandleCommand(interaction *discordgo.InteractionCreate) error {
	data := interaction.ApplicationCommandData()
	command, exists := bot.Commands[data.Name]
	if !exists {
		log.Panicf("Interaction with name %s is not implemented into the bot", data.Name)
	}

	if command.ListenInDM && !command.ListenInPublic && interaction.GuildID != "" {
		return bot.UserError(interaction, "🚫 Cette commande ne peut être executée qu'en message privé.")
	}
	if command.ListenInPublic && !command.ListenInDM && interaction.GuildID == "" {
		return bot.UserError(interaction, "🚫 Cette commande ne peut être executée que dans un salon public.")
	}

	return command.Execute(bot, interaction)
}

func (bot *Bot) HandleComponent(interaction *discordgo.InteractionCreate) error {
	data := interaction.MessageComponentData()
	component, exists := bot.Components[data.CustomID]
	if !exists {
		log.Panicf("Component with custom ID %s is not implemented into the bot", data.CustomID)
	}

	return component(bot, interaction)
}

func (bot *Bot) UserError(interaction *discordgo.InteractionCreate, message string) error {
	return bot.Client.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: 4,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				bot.MakeEmbed(&discordgo.MessageEmbed{
					Title: message,
					Color: bot.Config.Bot.ErrorColor,
				}),
			},
		},
	})
}

// MakeEmbed returns a Discord embed with the style of the bot
func (bot *Bot) MakeEmbed(base *discordgo.MessageEmbed) *discordgo.MessageEmbed {
	color := bot.Config.Bot.Color
	if base.Color != 0 {
		color = base.Color
	}
	return &discordgo.MessageEmbed{
		Title:       base.Title,
		Description: base.Description,
		Color:       color,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05-0700"),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("%s v%s", bot.Config.Bot.Name, bot.Config.Dev.Version),
			IconURL: bot.Config.Bot.Illustration,
		},
		Image:     base.Image,
		Thumbnail: base.Thumbnail,
		Video:     base.Video,
		Provider:  base.Provider,
		Author:    base.Author,
		Fields:    base.Fields,
	}
}
