package main

import (
	"strings"

	"github.com/theovidal/onyxcord/handlers"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.HasPrefix(message.Content, bot.Config.Bot.Prefix) && !message.Author.Bot {
		bot.OnCommand(session, message)
	}
}

func MessageDelete(_ *discordgo.Session, message *discordgo.MessageDelete) {
	handlers.ReactionRoleHandlerRemove(&bot, message)
}

func ReactionAdd(_ *discordgo.Session, reaction *discordgo.MessageReactionAdd) {
	handlers.ReactionRoleAdd(&bot, reaction)
}

func ReactionRemove(_ *discordgo.Session, reaction *discordgo.MessageReactionRemove) {
	handlers.ReactionRoleRemove(&bot, reaction)
}
