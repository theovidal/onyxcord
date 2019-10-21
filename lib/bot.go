package lib

import (
	"fmt"
	"log"
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

func (bot *Bot) OnCommand(session disgord.Session, context *disgord.MessageCreate) {
	msg := context.Message
	composition := strings.Split(msg.Content, " ")
	if commandName := composition[0]; strings.HasPrefix(commandName, bot.Config.Bot.Prefix) {
		command, exists := bot.Commands[strings.TrimPrefix(commandName, bot.Config.Bot.Prefix)]
		if exists {
			execution, success := command.Execute.(func(arguments []string, bot Bot, context *disgord.MessageCreate))
			if success {
				bot.Session = &session
				arguments := composition[1:]
				execution(arguments, *bot, context)
			} else {
				log.Fatalf("Unknown error while executing a command")
			}
		} else {
			_, _ = msg.Respond(
				session,
				&disgord.Message{
					Embeds: []*disgord.Embed{
						{
							Title: ":x: Erreur",
							Description: fmt.Sprintf("La commande %s est inconnue."+
								"\nFaites `%s@help` pour obtenir la liste des commandes disponibles.",
								commandName, bot.Config.Bot.Prefix),
							Color: bot.Config.Bot.Color,
						},
					},
				},
			)
		}
	}
}
