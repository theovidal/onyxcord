package cuatro

import (
	"github.com/bwmarrin/discordgo"
	"github.com/theovidal/onyxcord/bots/cuatro/handlers"

	"github.com/theovidal/onyxcord/bots/cuatro/commands"
	"github.com/theovidal/onyxcord/lib"
)

func Install() *lib.Bot {
	bot := lib.RegisterBot("cuatro", commands.List)
	bot.Client.AddHandler(func(_ *discordgo.Session, message *discordgo.MessageDelete) {
		handlers.ReactionRoleHandlerRemove(&bot, message)
	})
	bot.Client.AddHandler(func(_ *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
		handlers.ReactionRoleAdd(&bot, reaction)
	})
	bot.Client.AddHandler(func(_ *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
		handlers.ReactionRoleRemove(&bot, reaction)
	})
	bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions)
	return &bot
}
