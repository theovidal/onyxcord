package boite_a_bois

import (
	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/bots/boite_a_bois/commands"
	"github.com/theovidal/onyxcord/lib"
)

func Install() *lib.Bot {
	bot := lib.RegisterBot("boite_a_bois", commands.List)
	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	return &bot
}
