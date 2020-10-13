package carrbridge

import (
	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/bots/carrbridge/handlers"
	"github.com/theovidal/onyxcord/lib"
)

func Install() lib.Bot {
	bot := lib.RegisterBot("carrbridge")

	bot.Client.AddHandler(func(_ *discordgo.Session, message *discordgo.MessageCreate) {
		handlers.MessageTransfer(&bot, message)
	})

	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	return bot
}
