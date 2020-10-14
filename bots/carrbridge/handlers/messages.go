package handlers

import (
	"context"
	"encoding/hex"
	"log"
	"math/rand"

	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/theovidal/onyxcord/lib"
)

func _() string {
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

type Router struct {
	Token       string `json:"token"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Plugs       []Plug `json:"plugs"`
}

type Plug struct {
	Channel string `json:"channel"`
	Webhook string `json:"webhook"`
	Token   string `json:"token"`
}

func MessageTransfer(bot *lib.Bot, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	ctx := context.Background()
	routers := lib.Db.Database(bot.Config.Database).Collection(bot.Config.Assets["routers_collection"])

	var connectedRouters []Router
	results, err := routers.Find(ctx, bson.M{
		"plugs": bson.M{
			"$elemMatch": bson.M{
				"channel": msg.ChannelID,
			},
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	err = results.All(ctx, &connectedRouters)
	if err != nil {
		log.Println(err)
		return
	}

	for _, router := range connectedRouters {
		for _, plug := range router.Plugs {
			if plug.Channel == msg.ChannelID {
				continue
			}
			avatar := msg.Author.AvatarURL("128")
			var name string
			member, _ := bot.Client.GuildMember(msg.GuildID, msg.Author.ID)
			if member.Nick != "" {
				name = member.Nick
			} else {
				name = msg.Author.Username
			}

			content := msg.Content
			for _, file := range msg.Attachments {
				content += file.ProxyURL
			}

			if content != "" {
				_, err := bot.Client.WebhookExecute(plug.Webhook, plug.Token, false, &discordgo.WebhookParams{
					Username:  name,
					AvatarURL: avatar,
					Content:   content,
				})
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
