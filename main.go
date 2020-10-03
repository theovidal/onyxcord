package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/commands"
	"github.com/theovidal/onyxcord/lib"

	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
 * This file is the entry point of the bot, where it's setup.
 */

var bot = lib.LoadBot(commands.List)

func main() {
	bot.Client.AddHandler(MessageCreate)
	bot.Client.AddHandler(ReactionAdd)
	bot.Client.AddHandler(ReactionRemove)
	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions)

	// Open a websocket connection to Discord and begin listening.
	err := bot.Client.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Printf("Logged in as %s#%s\n", bot.User.Username, bot.User.Discriminator)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down sessions
	bot.Client.Close()
	bot.Db.Client.Disconnect(context.Background())
	log.Println("Goodbye")
}
