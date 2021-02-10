package onyxcord

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

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
	// A list of all the commands available on the bot
	Commands map[string]*Command
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
	bot.Commands["help"] = Help()
	return bot
}

func (bot *Bot) RegisterCommand(name string, command *Command) {
	bot.Commands[name] = command
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
		bot.Client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
			bot.OnCommand(session, message)
		})
	}

	log.Println("⏩ Connecting bot...")
	// Open a websocket connection to Discord and begin listening.
	err := bot.Client.Open()
	if err != nil {
		log.Fatalf("‼ Error opening connection with bot %s: %s\n", bot.Name, err)
	}
	log.Printf("✅ %s#%s logged in!\n", bot.User.Username, bot.User.Discriminator)

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

// OnCommand reacts to a newly-created message and treats it
func (bot *Bot) OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, bot.Config.Bot.Prefix) || message.Author.Bot {
		return
	}
	message.Content = strings.TrimPrefix(message.Content, bot.Config.Bot.Prefix)
	parts := strings.Split(message.Content, " ")
	commandName := parts[0]
	command, exists := bot.Commands[commandName]
	if !exists {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID,
			MakeEmbed(bot.Config, &discordgo.MessageEmbed{
				Title: fmt.Sprintf(":question: La commande %s est inconnue", commandName),
				Description: fmt.Sprintf(
					"\nExécutez `%shelp` pour obtenir la liste des commandes disponibles.",
					bot.Config.Bot.Prefix),
				Color: bot.Config.Bot.Color,
			}),
		)
		return
	}

	if !(command.ListenInDM && message.GuildID == "" || command.ListenInPublic && message.GuildID != "") {
		return
	}

	// TODO: permission check (UserChannelPermissions method always returns 0)
	/*userPermissions, _ := session.State.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if command.Permissions != 0 && command.Permissions & userPermissions == 0 {
		_, _ = session.ChannelMessageSendEmbed(message.ChannelID,
			MakeEmbed(bot.Config, &discordgo.MessageEmbed{
				Title: ":x: Vous n'avez pas la permission d'exécuter cette commande.",
				Color: 12000284,
			}),
		)
		return
	}*/

	argumentsPart := strings.Join(parts[1:], " ")
	arguments := strings.Split(argumentsPart, ",")
	bot.ExecuteCommand(command, arguments, message)
}

// ExecuteCommand executes the command parsed in the OnCommand function
// **It shouldn't be used by the end-user, but is stayed as public for flexibility**
func (bot *Bot) ExecuteCommand(command *Command, arguments []string, message *discordgo.MessageCreate) {
	err := command.Execute(arguments, bot, message)
	if err != nil {
		_, _ = bot.Client.ChannelMessageSendEmbed(message.ChannelID,
			MakeEmbed(bot.Config, &discordgo.MessageEmbed{
				Title: ":x: Erreur dans l'exécution de la commande",
				Description: fmt.Sprintf(
					"**%v**\n\n"+
						"N'hésitez-pas à contacter %s (%s) si vous pensez que c'est un bogue !",
					err, bot.Config.Dev.Maintainer.Name,
					bot.Config.Dev.Maintainer.Link,
				),
				Color: 12000284,
			}),
		)
	}
}
