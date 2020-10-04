package commands

import (
	"context"
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/bots/cuatro/handlers"
	"github.com/theovidal/onyxcord/lib"
)

func ReactionRole() *lib.Command {
	return &lib.Command{
		Description:    "Ajouter un support de réaction-rôle",
		Usage:          "reactionRole <message_link>",
		Category:       "utilities",
		Show:           false,
		ListenInPublic: true,
		Execute: func(arguments []string, bot lib.Bot, message *discordgo.MessageCreate) (err error) {
			if len(arguments) == 0 {
				return errors.New("Vous devez joindre un lien vers un message")
			}

			parts := strings.Split(arguments[0], "/")
			channelID := strings.Trim(parts[5], " ")
			messageID := strings.Trim(parts[6], " ")
			support, err := bot.Client.ChannelMessage(
				channelID,
				messageID,
			)
			if err != nil {
				return errors.New(
					"Le lien vers le message est invalide." +
						"Vérifiez qu'il existe encore et/ou que le robot a la permission d'accéder au salon souhaité.",
				)
			}

			roles := make(map[string]string)
			for i, part := range arguments[1:] {
				emoji := part
				bot.Client.MessageReactionAdd(support.ChannelID, support.ID, emoji)
				roles[emoji] = support.MentionRoles[i]
			}

			_, err = lib.Db.Database(bot.Config.Database).
				Collection("reactionRoles").
				InsertOne(context.Background(), handlers.ReactionRole{
					Message: messageID,
					Channel: channelID,
					Roles:   roles,
				})

			_, _ = bot.Client.ChannelMessageSendEmbed(
				message.ChannelID,
				lib.MakeEmbed(
					bot.Config,
					&discordgo.MessageEmbed{
						Title:       ":white_check_mark: Le support de réaction-rôle a été ajouté.",
						Description: "Supprimez le message pour désactiver le réaction-rôle sur celui-ci",
					},
				),
			)
			return
		},
	}
}
