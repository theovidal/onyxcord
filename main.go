package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/BecauseOfProg/boite-a-bois/commands"
	"github.com/BecauseOfProg/boite-a-bois/lib"

	"github.com/andersfylling/disgord"
)

/*
 * This file is the entry point of the bot, where it's setup.
 * For the moment is just a bunch of code that'll, in the future, be split between multiple files.
 */

func main() {
	config, err := lib.GetConfig(filepath.Base("./config.yml"))
	if err != nil {
		log.Fatalf("Error while getting the configuration : %v", err)
	}

	discord := disgord.New(&disgord.Config{
		BotToken: config.Bot.Token,
	})
	defer discord.StayConnectedUntilInterrupted()

	user, err := discord.GetUser(
		disgord.NewSnowflake(
			uint64(config.Bot.ID),
		),
	)
	if err != nil {
		log.Fatalf("Error while creating the user profile : %v", err)
		return
	}

	bot := lib.Bot{
		Commands: commands.List,
		Config:   &config,
		User:     user,
	}

	// Listen for incoming messages and parse them as commands
	discord.On(disgord.EvtMessageCreate, func(session disgord.Session, context *disgord.MessageCreate) {
		msg := context.Message
		composition := strings.Split(msg.Content, " ")
		if commandName := composition[0]; strings.HasPrefix(commandName, config.Bot.Prefix) {
			command, exists := bot.Commands[strings.TrimPrefix(commandName, config.Bot.Prefix)]
			if exists {
				execution, success := command.Execute.(func(arguments []string, bot lib.Bot, context *disgord.MessageCreate))
				if success {
					bot.Session = &session
					arguments := composition[1:]
					execution(arguments, bot, context)
				} else {
					log.Fatalf("Error while executing a command : %v", err)
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
									commandName, config.Bot.Prefix),
								Color: config.Bot.Color,
							},
						},
					},
				)
			}
		}
	})

	discord.Ready(func() {
		fmt.Println("Successfully logged in.")
	})
}
