package onyxcord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Help() *Command {
	return &Command{
		Description:    "Obtenir de l'aide sur les commandes du robot",
		Show:           true,
		ListenInDM:     true,
		ListenInPublic: true,
		Execute: func(arguments []string, bot Bot, message *discordgo.MessageCreate) (err error) {
			commandsList := make(map[string]string)
			for name, command := range bot.Commands {
				if !command.Show {
					continue
				} else {
					if command.Category == "" {
						command.Category = "default"
					}
					commandsList[command.Category] += command.Prettify(name, bot.Config.Bot.Prefix) + "\n"
				}
			}

			fullMessage := discordgo.MessageEmbed{
				Title: ":fast_forward: Commandes du robot",
				Description: fmt.Sprintf(
					"Voici une liste des commandes disponibles sur %s. "+
						"Elles s'exécutent dans un salon textuel avec, ou non, des arguments séparés par une virgule.",
					bot.Config.Bot.Name,
				),
			}
			for categoryName, commands := range commandsList {
				category := bot.Config.Categories[categoryName]
				fullMessage.Fields = append(fullMessage.Fields, &discordgo.MessageEmbedField{
					Name:   fmt.Sprintf(":%s: %s", category.Emoji, category.Name),
					Value:  commands,
					Inline: true,
				})
			}

			fmt.Println(message.ChannelID, message.Author.ID)
			_, _ = bot.Client.ChannelMessageSendEmbed(message.ChannelID, MakeEmbed(bot.Config, &fullMessage))
			return
		},
	}
}
