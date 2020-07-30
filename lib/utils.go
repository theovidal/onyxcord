package lib

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"time"
)

// MakeEmbed returns a Discord embed with the style of the bot
func MakeEmbed(config *Config, base *disgord.Embed) *disgord.Embed {
	return &disgord.Embed{
		Title:       base.Title,
		Description: base.Description,
		Timestamp:   disgord.Time{Time: time.Now().UTC()},
		Color:       config.Bot.Color,
		Footer: &disgord.EmbedFooter{
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
