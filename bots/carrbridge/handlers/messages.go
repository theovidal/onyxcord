package handlers

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/theovidal/onyxcord/lib"
)

type Router struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Description string
}

type Plug struct {
	ID      uint `gorm:"primary_key"`
	Router  uint
	Channel string
	Webhook string
	Token   string
}

func MessageTransfer(bot *lib.Bot, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}

	filename := "testing.sqlite"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Error while creating the database: %s", err)
		}
	}

	db, err := gorm.Open("sqlite3", filename)
	if err != nil {
		log.Fatalf("Failed to open database: %s", err)
	}

	var originPlug Plug
	db.Where("channel = ?", msg.ChannelID).First(&originPlug)
	if originPlug.Router == 0 {
		return
	}

	var plugs []Plug
	db.Where("router = ?", originPlug.Router).Find(&plugs)
	for _, plug := range plugs {
		if plug.ID == originPlug.ID {
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

		_, err := bot.Client.WebhookExecute(plug.Webhook, plug.Token, false, &discordgo.WebhookParams{
			Content:   msg.Content,
			Username:  name,
			AvatarURL: avatar,
		})
		if err != nil {
			log.Println(err)
		}
	}
}
