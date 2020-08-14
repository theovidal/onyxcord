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
	Channel int                `bson:"channel"`
	Message int                `bson:"message"`
	Roles   map[string]int     `bson:"roles"`
}

func (bot *Bot) OnReactionAdd(_ disgord.Session, context *disgord.MessageReactionAdd) {
	found, reaction := FindReactionRole(bot, &context.ChannelID, &context.MessageID)
	if !found {
		return
	}

	channel, _ := bot.Client.GetChannel(context.Ctx, context.ChannelID)
	user, _ := bot.Client.GetMember(context.Ctx, channel.GuildID, context.UserID)

	role := reaction.Roles[context.PartialEmoji.Name]
	println(role)
	println(743574037992702032)
	println(role == 743574037992702032)
	userRoles := append(user.Roles, disgord.NewSnowflake(uint64(role)))

	err := bot.Client.UpdateGuildMember(context.Ctx, channel.GuildID, context.UserID).SetRoles(userRoles).Execute()
	if err != nil {
		fmt.Println(err)
	}
}

func (bot *Bot) OnReactionRemove(_ disgord.Session, _ *disgord.MessageReactionRemove) {

}

func FindReactionRole(bot *Bot, _ *disgord.Snowflake, _ *disgord.Snowflake) (found bool, data ReactionRole) {
	collection := bot.Db.Database(bot.Config.Database.Database).Collection("reactionRoles")

	filter := bson.M{}

	err := collection.FindOne(context.TODO(), filter).Decode(&data)

	fmt.Println(data)

	if err == nil {
		found = true
	}
	return
}
