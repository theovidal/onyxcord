package onyxcord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Command represents a command that can be executed by a user
type Command struct {
	// Description of the command (e.g : Prints weather for a specific location)
	Description string
	// Usage of the command (e.g : weather <location>)
	Usage string
	// Category of the command, as defined in the configuration (e.g : utilities)
	Category string
	// An alias to another command (e.g : w)
	Alias string
	// Choose if the command is shown in the help or not
	Show bool
	// Whether the bot should listen to the command in public channels
	ListenInPublic bool
	// Whether the bot should listen to the command in direct messages
	ListenInDM bool
	// Required permissions for the user to execute the command
	// See https://discordapi.com/permissions.html to generate the permission integer you want
	Permissions int
	// Action to execute if the command is triggered
	Execute func(arguments []string, bot *Bot, message *discordgo.MessageCreate) (err error)
}

// Prettify returns a string with information about a command, ready to be printed to the user
func (command Command) Prettify(name string, prefix string) (prettified string) {
	prettified = fmt.Sprintf("● **%s%s** : %s", prefix, name, command.Description)
	if command.Usage != "" {
		prettified += fmt.Sprintf("\nUtilisation : `%s`", command.Usage)
	}
	if command.Alias != "" {
		prettified += fmt.Sprintf("\nAlias : `%s`", command.Alias)
	}
	if !(command.ListenInPublic && command.ListenInDM) {
		prettified += "\nS'exécute uniquement dans "
		if command.ListenInDM {
			prettified += "les messages privés"
		} else {
			prettified += "un salon du serveur"
		}
	}
	return
}
