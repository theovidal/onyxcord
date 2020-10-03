package main

import (
	"context"
	"fmt"
	"github.com/theovidal/onyxcord/commands"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/theovidal/onyxcord/bots"
	"github.com/theovidal/onyxcord/lib"

	_ "go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
 * This file is the entry point of the bot, where it's setup.
 */

func main() {
	lib.OpenDatabase()

	for bot := range bots.Bots {
		bot.Commands["help"] = &commands.Help
		// Open a websocket connection to Discord and begin listening.
		err := bot.Client.Open()
		if err != nil {
			fmt.Println("error opening connection with bot: ", err)
			return
		}
		fmt.Println(bot)
		log.Printf("%s#%s logged in\n", bot.User.Username, bot.User.Discriminator)
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Println("Closing down...")

	for bot := range bots.Bots {
		// Cleanly close down sessions
		bot.Client.Close()
		log.Printf("%s#%s disconnected\n", bot.User.Username, bot.User.Discriminator)
	}
	lib.Db.Disconnect(context.Background())
	log.Println("Goodbye")
}
