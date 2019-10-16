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

var bot = lib.Bot{
	Commands: &[]lib.Command{},
	Config:   &lib.Config{},
}

func main() {
	config, err := lib.GetConfig(filepath.Base("./config.yml"))
	if err != nil {
		log.Fatalf("Error while getting the configuration : %v", err)
	}

	discord := disgord.New(&disgord.Config{
		BotToken: config.Bot.Token,
	})
	defer discord.StayConnectedUntilInterrupted()

	bot = lib.Bot{
		Commands: &[]lib.Command{},
		Config:   &config,
	}

	// Listen for incoming messages and parse them as commands
	discord.On(disgord.EvtMessageCreate, func(session disgord.Session, context *disgord.MessageCreate) {
		bot.Session = &session
		msg := context.Message
		composition := strings.Split(msg.Content, " ")
		command := composition[0]
		arguments := composition[1:]
		if command == "bd@ping" {
			_, err = msg.Reply(session, "pong")
			if err != nil {
				log.Fatalf("Error while sending a message : %v", err)
			}
		}
		if command == "bd@weather" {
			commands.Weather.Execute.(func(arguments []string, bot lib.Bot, context *disgord.MessageCreate))(arguments, bot, context)
		}
	})

	discord.Ready(func() {
		fmt.Println("Successfully logged in.")
	})
}
