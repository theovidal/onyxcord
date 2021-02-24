package onyxcord

import (
	"github.com/bwmarrin/discordgo"
)

// Command represents a command that can be executed by a user
type Command struct {
	// Whether the bot should listen to the command in public channels
	ListenInPublic bool
	// Whether the bot should listen to the command in direct messages
	ListenInDM bool
	// Required permissions for the user to execute the command
	// See https://discordapi.com/permissions.html to generate the permission integer you want
	Permissions int
	// Action to execute if the command is triggered
	Execute func(bot *Bot, interaction *discordgo.InteractionCreate) (err error)
}
