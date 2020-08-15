package commands

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/theovidal/onyxcord/lib"
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

		fullMessage := disgord.Embed{
			Title: ":fast_forward: Commandes du robot",
			Description: fmt.Sprintf(
				"Voici une liste des commandes disponibles sur %s."+
					"Elles s'exécutent dans un salon textuel avec, ou non, des arguments séparés par une virgule.",
				bot.Config.Bot.Name,
			),
		}
		for categoryName, commands := range commandsList {
			category := bot.Config.Categories[categoryName]
			fullMessage.Fields = append(fullMessage.Fields, &disgord.EmbedField{
				Name:   fmt.Sprintf(":%s: %s", category.Emoji, category.Name),
				Value:  commands,
				Inline: true,
			})
		}

		_ = bot.SendEmbed(
			context.Ctx,
			&fullMessage,
			context.Message,
		)
		return
	},
}
