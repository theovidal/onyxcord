package handlers

import (
	"context"
	"fmt"

	"github.com/andersfylling/disgord"
	"github.com/theovidal/onyxcord/lib"
	"go.mongodb.org/mongo-driver/bson"
)

type ReactionRole struct {
	Channel string
	Message string
	Roles   map[string]string
}

func remove(s []disgord.Snowflake, i int) []disgord.Snowflake {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func ReactionRoleAdd(bot *lib.Bot, context *disgord.MessageReactionAdd) {
	if context.UserID == bot.User.ID {
		return
	}
	found, role, channel, userRoles := FindReactionRole(bot, context.ChannelID, context.MessageID, context.UserID, context.PartialEmoji.Name)
	if !found {
		return
	}

	userRoles = append(userRoles, role)
	SetRoles(bot, channel.GuildID, context.UserID, userRoles)
}

func ReactionRoleRemove(bot *lib.Bot, context *disgord.MessageReactionRemove) {
	if context.UserID == bot.User.ID {
		return
	}
	found, role, channel, userRoles := FindReactionRole(bot, context.ChannelID, context.MessageID, context.UserID, context.PartialEmoji.Name)
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

func ReactionRoleHandlerRemove(bot *lib.Bot, context *disgord.MessageDelete) {
	filter := GetFilter(context.ChannelID, context.MessageID)
	result, _ := bot.Db.ReactionRoles.DeleteOne(context.Ctx, filter)
	if result.DeletedCount == 0 {
		return
	}

	_, _ = bot.Client.SendMsg(
		context.Ctx,
		context.ChannelID,
		&disgord.CreateMessageParams{
			Embed: lib.MakeEmbed(
				bot.Config,
				&disgord.Embed{
					Title:       ":white_check_mark: Le support de réaction-rôle a été supprimé.",
					Description: "Si vous pensez que c'est une erreur, n'hésitez-pas à nous le signaler!",
				},
			),
		},
	)
}

func GetFilter(channelID, messageID disgord.Snowflake) bson.M {
	return bson.M{
		"channel": channelID.String(),
		"message": messageID.String(),
	}
}

func FindReactionRoleHandler(bot *lib.Bot, channelID, messageID disgord.Snowflake) (found bool, channel *disgord.Channel, data ReactionRole) {
	filter := GetFilter(channelID, messageID)
	err := bot.Db.ReactionRoles.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return
	}

	found = true
	channel, _ = bot.Client.GetChannel(context.Background(), channelID)
	return
}

func FindReactionRole(bot *lib.Bot, channelID, messageID, userID disgord.Snowflake, emoji string) (found bool, role disgord.Snowflake, channel *disgord.Channel, userRoles []disgord.Snowflake) {
	found, channel, data := FindReactionRoleHandler(bot, channelID, messageID)
	if !found {
		return
	}

	role = disgord.ParseSnowflakeString(data.Roles[emoji])
	user, _ := bot.Client.GetMember(context.Background(), channel.GuildID, userID)
	userRoles = user.Roles
	return
}

func SetRoles(bot *lib.Bot, guild, user disgord.Snowflake, roles []disgord.Snowflake) {
	err := bot.Client.UpdateGuildMember(context.Background(), guild, user).SetRoles(roles).Execute()
	if err != nil {
		fmt.Println(err)
	}
}
