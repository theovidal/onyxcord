package onyxcord

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bwmarrin/discordgo"
)

// OpenFile opens a file from a path
func OpenFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	return
}

// MakeEmbed returns a Discord embed with the style of the bot
func MakeEmbed(config *Config, base *discordgo.MessageEmbed) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       base.Title,
		Description: base.Description,
		Color:       config.Bot.Color,
		Timestamp:   time.Now().Format("2006-01-02T15:04:05-0700"),
		Footer: &discordgo.MessageEmbedFooter{
			Text:    fmt.Sprintf("%s v%s", config.Bot.Name, config.Dev.Version),
			IconURL: config.Bot.Illustration,
		},
		Image:     base.Image,
		Thumbnail: base.Thumbnail,
		Video:     base.Video,
		Provider:  base.Provider,
		Author:    base.Author,
		Fields:    base.Fields,
	}
}
