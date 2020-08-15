package handlers

import (
	"context"
	"fmt"

	"github.com/BecauseOfProg/boite-a-bois/lib"
	"github.com/andersfylling/disgord"
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

func FindReactionRole(bot *lib.Bot, channelID, messageID, userID disgord.Snowflake, emoji string) (found bool, role disgord.Snowflake, channel *disgord.Channel, userRoles []disgord.Snowflake) {
	filter := bson.M{
		"channel": channelID.String(),
		"message": messageID.String(),
	}

	var data ReactionRole
	err := bot.Db.ReactionRoles.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return
	}

	found = true
	role = disgord.ParseSnowflakeString(data.Roles[emoji])
	channel, _ = bot.Client.GetChannel(context.Background(), channelID)
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
