package main

import (
	"context"
	"github.com/andersfylling/disgord/std"
	"log"

	"github.com/BecauseOfProg/boite-a-bois/commands"
	"github.com/BecauseOfProg/boite-a-bois/lib"
	"github.com/andersfylling/disgord"
)

/*
 * This file is the entry point of the bot, where it's setup.
 */

func main() {
	bot := lib.LoadBot(commands.List)
	defer bot.Client.StayConnectedUntilInterrupted(context.Background())

	commandFilter, _ := std.NewMsgFilter(context.Background(), bot.Client)
	commandFilter.SetPrefix(bot.Config.Bot.Prefix)

	// Listen for incoming messages and parse them as commands
	bot.Client.On(disgord.EvtMessageCreate, commandFilter.HasPrefix, func(session disgord.Session, context *disgord.MessageCreate) {
		bot.OnCommand(session, context)
	})

	// Print to the console when the bot is ready
	bot.Client.Ready(func() {
		log.Printf("Logged in as %s#%d\n", bot.User.Username, bot.User.Discriminator)
	})
}
