package lib

import "github.com/bwmarrin/discordgo"

type SkeletonBot struct {
	// The name of the bot to create
	Name string
	// A list of all the commands available on the bot
	Commands map[string]*Command

	Handlers []func(session *discordgo.Session, event interface{})
}
