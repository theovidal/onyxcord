package main

import (
	"fmt"
	"github.com/BecauseOfProg/boite-a-bois/commands"
	"github.com/BecauseOfProg/boite-a-bois/lib"
	"log"
	"path/filepath"

	"github.com/andersfylling/disgord"
)

/*
 * This file is the entry point of the bot, where it's setup.
 * For the moment is just a bunch of code that'll, in the future, be split between multiple files.
 */

func main() {
	// Loading the configuration
	config, err := lib.GetConfig(filepath.Base("./config.yml"))
	if err != nil {
		log.Panicf("Error while getting the configuration : %v", err)
	}

	// Creating the bot profile
	client := disgord.New(&disgord.Config{
		BotToken: config.Bot.Token,
		Presence: &disgord.UpdateStatusCommand{
			Since:  nil,
			Game:   nil,
			Status: fmt.Sprintf("%shelp", config.Bot.Prefix),
			AFK:    false,
		},
	})
	defer client.StayConnectedUntilInterrupted()

	user, err := client.GetUser(
		disgord.NewSnowflake(
			uint64(config.Bot.ID),
		),
	)
	if err != nil {
		log.Panicf("Error while creating the user profile : %v", err)
	}

	bot := lib.Bot{
		Client:   client,
		Commands: commands.List,
		Config:   &config,
		User:     user,
	}

	// Listen for incoming messages and parse them as commands
	client.On(disgord.EvtMessageCreate, func(session disgord.Session, context *disgord.MessageCreate) {
		bot.OnCommand(session, context)
	})

	client.Ready(func() {
		log.Printf("Successfully logged in as %s#%d\n", bot.User.Username, bot.User.Discriminator)
	})
}
