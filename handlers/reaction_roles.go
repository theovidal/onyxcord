package handlers

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/theovidal/onyxcord/lib"
)

type ReactionRole struct {
	Channel string
	Message string
	Roles   map[string]string
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func ReactionRoleAdd(bot *lib.Bot, context *discordgo.MessageReactionAdd) {
	if context.UserID == bot.User.ID {
		return
	}
	found, role, channel, userRoles := FindReactionRole(bot, context.ChannelID, context.MessageID, context.UserID, context.Emoji.Name)
	if !found {
		return
	}

	userRoles = append(userRoles, role)
	SetRoles(bot, channel.GuildID, context.UserID, userRoles)
}

func ReactionRoleRemove(bot *lib.Bot, context *discordgo.MessageReactionRemove) {
	if context.UserID == bot.User.ID {
		return
	}
	found, role, channel, userRoles := FindReactionRole(bot, context.ChannelID, context.MessageID, context.UserID, context.Emoji.Name)
	if !found {
		return
	}

	var indexToRemove int
	for index, roleID := range userRoles {
		if roleID == role {
			indexToRemove = index
			break
		}
	}

	userRoles = remove(userRoles, indexToRemove)
	SetRoles(bot, channel.GuildID, context.UserID, userRoles)
}

func ReactionRoleHandlerRemove(bot *lib.Bot, message *discordgo.MessageDelete) {
	filter := GetFilter(message.ChannelID, message.ID)
	result, _ := bot.Db.ReactionRoles.DeleteOne(context.Background(), filter)
	if result.DeletedCount == 0 {
		return
	}

	_, _ = bot.Client.ChannelMessageSendEmbed(
		message.ChannelID,
		lib.MakeEmbed(
			bot.Config,
			&discordgo.MessageEmbed{
				Title:       ":white_check_mark: Le support de réaction-rôle a été supprimé.",
				Description: "Si vous pensez que c'est une erreur, n'hésitez-pas à nous le signaler!",
			},
		),
	)
}

func GetFilter(channelID, messageID string) bson.M {
	return bson.M{
		"channel": channelID,
		"message": messageID,
	}
}

func FindReactionRoleHandler(bot *lib.Bot, channelID, messageID string) (found bool, channel *discordgo.Channel, data ReactionRole) {
	filter := GetFilter(channelID, messageID)
	err := bot.Db.ReactionRoles.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		return
	}

	found = true
	channel, _ = bot.Client.Channel(channelID)
	return
}

func FindReactionRole(bot *lib.Bot, channelID, messageID, userID string, emoji string) (found bool, role string, channel *discordgo.Channel, userRoles []string) {
	found, channel, data := FindReactionRoleHandler(bot, channelID, messageID)
	if !found {
		return
	}

	role = data.Roles[emoji]
	user, _ := bot.Client.GuildMember(channel.GuildID, userID)
	userRoles = user.Roles
	return
}

func SetRoles(bot *lib.Bot, guild, user string, roles []string) {
	err := bot.Client.GuildMemberEdit(guild, user, roles)
	if err != nil {
		fmt.Println(err)
	}
}
