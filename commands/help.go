package commands

import (
	"fmt"
	"github.com/BecauseOfProg/boite-a-bois/lib"
	"github.com/andersfylling/disgord"
)

var help = lib.Command{
	Description: "Obtenir de l'aide sur les commandes du robot",
	Usage:       "help",
	Category:    "utilities",
	Listen:      []string{"public", "private"},
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		commandsList := make(map[string]string)
		for name, command := range bot.Commands {
			if command.Show {
				continue
			} else {
				commandsList[command.Category] += command.Prettify(name, bot.Config.Bot.Prefix) + "\n"
			}
		}

		var fullMessage disgord.Embed
		for categoryName, commands := range commandsList {
			category := bot.Config.Categories[categoryName]
			fullMessage.Fields = append(fullMessage.Fields, &disgord.EmbedField{
				Name:  fmt.Sprintf(":%s: %s", category.Emoji, category.Name),
				Value: commands,
			})
		}

		_ = bot.SendEmbed(
			&fullMessage,
			context.Message,
		)
		return
	},
}
