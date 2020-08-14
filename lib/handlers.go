package lib

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReactionRole struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Channel string             `bson:"channel"`
	Message string             `bson:"message"`
	Roles   map[string]string  `bson:"roles"`
}

func remove(s []disgord.Snowflake, i int) []disgord.Snowflake {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (bot *Bot) OnReactionAdd(_ disgord.Session, context *disgord.MessageReactionAdd) {
	found, role, channel, user := FindReactionRole(bot, &context.ChannelID, &context.MessageID, context.UserID, context.PartialEmoji.Name)
	if !found {
		return
	}

	userRoles := append(user.Roles, role)

	err := bot.Client.UpdateGuildMember(context.Ctx, channel.GuildID, context.UserID).SetRoles(userRoles).Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func (bot *Bot) OnReactionRemove(_ disgord.Session, context *disgord.MessageReactionRemove) {
	found, role, channel, user := FindReactionRole(bot, &context.ChannelID, &context.MessageID, context.UserID, context.PartialEmoji.Name)
	if !found {
		return
	}

	var indexToRemove int
	for index, roleID := range user.Roles {
		if roleID == role {
			indexToRemove = index
			break
		}
	}

	userRoles := remove(user.Roles, indexToRemove)

	err := bot.Client.UpdateGuildMember(context.Ctx, channel.GuildID, context.UserID).SetRoles(userRoles).Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func FindReactionRole(bot *Bot, channelID *disgord.Snowflake, messageID *disgord.Snowflake, userID disgord.Snowflake, emoji string) (found bool, role disgord.Snowflake, channel *disgord.Channel, user *disgord.Member) {
	collection := bot.Db.Database(bot.Config.Database.Database).Collection("reactionRoles")

	filter := bson.M{
		"channel": channelID.String(),
		"message": messageID.String(),
	}

	var data ReactionRole
	err := collection.FindOne(context.TODO(), filter).Decode(&data)
	if err == nil {
		found = true
	}

	channel, _ = bot.Client.GetChannel(context.Background(), *channelID)
	user, _ = bot.Client.GetMember(context.Background(), channel.GuildID, userID)
	role = disgord.ParseSnowflakeString(data.Roles[emoji])

	return
}
