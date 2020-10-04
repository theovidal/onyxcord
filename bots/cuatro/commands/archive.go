package commands

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/lib"
)

func Archive() *lib.Command {
	return &lib.Command{
		Description:    "Archiver les messages du salon",
		Usage:          "archive",
		Category:       "utilities",
		Show:           false,
		ListenInPublic: true,
		Execute: func(arguments []string, bot lib.Bot, message *discordgo.MessageCreate) (err error) {
			bot.Client.ChannelMessageSend(message.ChannelID, "Ouverture du fichier")
			var file *os.File
			file, err = os.OpenFile("./tmp/archive.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				return
			}

			var channel string
			var before string
			for {
				bot.Client.ChannelMessageSend(message.ChannelID, "Récupération de messages...")
				messages, err := bot.Client.ChannelMessages(message.ChannelID, 100, before, "", "")
				if err != nil {
					return err
				}
				if len(messages) <= 1 {
					break
				}

				for _, message := range messages {
					date, _ := message.Timestamp.Parse()
					channel = fmt.Sprintf(
						"%s%d/%d/%d %d:%d %s : %s\n",
						channel,
						date.Day(),
						date.Month(),
						date.Year(),
						date.Hour(),
						date.Minute(),
						message.Author.Username,
						message.Content,
					)
					before = message.ID
				}
			}

			bot.Client.ChannelMessageSend(message.ChannelID, "Écriture du fichier")
			_, err = file.Write([]byte(channel))
			bot.Client.ChannelMessageSend(message.ChannelID, "Done!")
			return
		},
	}
}
