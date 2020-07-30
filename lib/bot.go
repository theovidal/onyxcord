package lib

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/andersfylling/disgord"
)

// Bot represents the Discord bot with its assets
type Bot struct {
	// The Discord client associated with the bot
	Client *disgord.Client
	// A list of all the commands available on the bot
	Commands map[string]*Command
	// The configuration of the bot, as defined in the corresponding file
	Config *Config
	// The Discord session
	Session *disgord.Session
	// The profile of the bot
	User *disgord.User
}

// LoadBot creates a new instance of the Discord Bot
func LoadBot(commands map[string]*Command) Bot {
	// Loading the configuration
	config, err := GetConfig(filepath.Base("./config.yml"))
	if err != nil {
		log.Panicf("Error while getting the configuration : %v", err)
	}

	// Creating the bot profile
	client := disgord.New(disgord.Config{
		BotToken: config.Bot.Token,
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: fmt.Sprintf("%shelp", config.Bot.Prefix),
				Type: 2,
			},
		},
	})

	// Creating the user profile
	user, err := client.GetUser(
		context.Background(),
		disgord.NewSnowflake(
			uint64(config.Bot.ID),
		),
	)
	if err != nil {
		log.Panicf("Error while creating the user profile : %v", err)
	}

	return Bot{
		Client:   client,
		Commands: commands,
		Config:   &config,
		User:     user,
	}
}

// OnCommand reacts to a newly-created message and treats it
func (bot *Bot) OnCommand(session disgord.Session, context *disgord.MessageCreate) {
	bot.Session = &session
	msg := context.Message
	parts := strings.Split(msg.Content, " ")
	commandName := strings.TrimPrefix(parts[0], bot.Config.Bot.Prefix)
	command, exists := bot.Commands[commandName]
	if exists {
		argumentsPart := strings.Join(parts[1:], " ")
		arguments := strings.Split(argumentsPart, ",")
		bot.ExecuteCommand(command, arguments, context)
	} else {
		_, _ = context.Message.Reply(
			context.Ctx,
			bot.Client,
			&disgord.CreateMessageParams{
				Embed: MakeEmbed(bot.Config, &disgord.Embed{
					Title: fmt.Sprintf("La commande %s est inconnue", commandName),
					Description: fmt.Sprintf(
						"\nExécutez `%shelp` pour obtenir la liste des commandes disponibles.",
						bot.Config.Bot.Prefix),
					Color: bot.Config.Bot.Color,
				}),
			},
		)
	}
}

// ExecuteCommand executes the command parsed in the OnCommand function
func (bot *Bot) ExecuteCommand(command *Command, arguments []string, context *disgord.MessageCreate) {
	commandFunction, success :=
		command.Execute.(func(arguments []string, bot Bot, context *disgord.MessageCreate) (err error))

	if success {
		err := commandFunction(arguments, *bot, context)
		if err != nil {
			_, _ = context.Message.Reply(
				context.Ctx,
				bot.Client,
				&disgord.CreateMessageParams{
					Embed: MakeEmbed(bot.Config, &disgord.Embed{
						Title: ":x: Erreur dans l'exécution de la commande",
						Description: fmt.Sprintf(
							"**%v**\n\n"+
								"N'hésitez-pas à contacter %s (%s) si vous pensez que c'est un bogue !",
							err, bot.Config.Dev.Maintainer.Name,
							bot.Config.Dev.Maintainer.Link,
						),
						Color: 12000284,
					}),
				},
			)
			log.Printf("Command %s (executed by %s#%d) failed to execute : %v",
				command.Usage, context.Message.Author.Username, context.Message.Author.Discriminator, err)
		}
	} else {
		log.Panicf("Unknown error while parsing a command")
	}
}

// SendEmbed is a helper that sends an embed in the current channel
func (bot Bot) SendEmbed(context context.Context, embed *disgord.Embed, msg *disgord.Message) (err error) {
	_, err = msg.Reply(
		context,
		bot.Client,
		&disgord.CreateMessageParams{
			Embed: MakeEmbed(
				bot.Config,
				embed,
			),
		},
	)
	return
}
