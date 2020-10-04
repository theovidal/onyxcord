package lib

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Bot represents the Discord bot with its assets
type Bot struct {
	// Name of the bot
	Name string
	// The Discord client associated with the bot
	Client *discordgo.Session
	// A list of all the commands available on the bot
	Commands map[string]*Command
	// The configuration of the bot, as defined in the corresponding file
	Config *Config
	// The profile of the bot
	User *discordgo.User
}

// RegisterBot creates a new instance of the Discord Bot
func RegisterBot(name string) Bot {
	// Loading the configuration
	config, err := GetConfig(name)
	if err != nil {
		log.Panicf("Error while getting the configuration : %v", err)
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
		log.Panicf("Error while creating the user profile : %v", err)
	}

	bot := Bot{
		Name:     name,
		Client:   client,
		Commands: make(map[string]*Command),
		Config:   &config,
		User:     user,
	}
	client.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		bot.OnCommand(session, message)
	})
	return bot
}

func (bot *Bot) RegisterCommand(name string, command *Command) {
	bot.Commands[name] = command
}

// OnCommand reacts to a newly-created message and treats it
func (bot *Bot) OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {
	if !strings.HasPrefix(message.Content, bot.Config.Bot.Prefix) || message.Author.Bot {
		return
	}
	parts := strings.Split(message.Content, " ")
	commandName := strings.TrimPrefix(parts[0], bot.Config.Bot.Prefix)
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

	if command.ListenInDM && message.ChannelID == "" || command.ListenInPublic && message.ChannelID != "" {
		argumentsPart := strings.Join(parts[1:], " ")
		arguments := strings.Split(argumentsPart, ",")
		bot.ExecuteCommand(command, arguments, message)
	}
}

// ExecuteCommand executes the command parsed in the OnCommand function
func (bot *Bot) ExecuteCommand(command *Command, arguments []string, message *discordgo.MessageCreate) {
	err := command.Execute(arguments, *bot, message)
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
