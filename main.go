package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/BecauseOfProg/boite-a-bois/commands"
	"github.com/BecauseOfProg/boite-a-bois/lib"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
 * This file is the entry point of the bot, where it's setup.
 */

func main() {
	// Loading the configuration
	config, err := lib.GetConfig(filepath.Base("./config.yml"))
	if err != nil {
		log.Panicf("Error while getting the configuration : %v", err)
	}

	// Loading the database
	uri := fmt.Sprint("mongodb://", config.Database.Address, ":", config.Database.Port)
	client, err := mongo.NewClient(
		options.Client().ApplyURI(uri).SetAuth(options.Credential{
			Username:   config.Database.Username,
			Password:   config.Database.Password,
			AuthSource: config.Database.AuthSource,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Creating the bot instance
	bot := lib.LoadBot(&config, commands.List, client)
	defer bot.Client.StayConnectedUntilInterrupted(context.Background())

	// Listen for incoming messages and parse them as commands
	commandFilter, _ := std.NewMsgFilter(context.Background(), bot.Client)
	commandFilter.SetPrefix(bot.Config.Bot.Prefix)
	bot.Client.On(disgord.EvtMessageCreate, commandFilter.HasPrefix, func(session disgord.Session, context *disgord.MessageCreate) {
		bot.OnCommand(session, context)
	})

	// Print to the console when the bot is ready
	bot.Client.Ready(func() {
		log.Printf("Logged in as %s#%d\n", bot.User.Username, bot.User.Discriminator)
	})
}
