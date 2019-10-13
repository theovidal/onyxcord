package lib

import "github.com/andersfylling/disgord"

// Bot represents the Discord bot with its assets
type Bot struct {
	// A list of all the commands available on the bot
	Commands *[]Command
	// The configuration of the bot, as defined in the corresponding file
	Config *Config
	// The Discord session
	Session *disgord.Session
}