package commands

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"os"

	"github.com/theovidal/onyxcord/lib"
)

var archive = lib.Command{
	Description:    "Archiver les messages du salon",
	Usage:          "archive",
	Category:       "utils",
	Show:           false,
	ListenInPublic: true,
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		context.Message.Reply(context.Ctx, bot.Client, "Ouverture du fichier")
		var file *os.File
		file, err = os.OpenFile("tmp/archive.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			return
		}

		var channel string
		var before disgord.Snowflake
		for {
			context.Message.Reply(context.Ctx, bot.Client, "Récupération de messages...")
			messages, err := bot.Client.GetMessages(context.Ctx, context.Message.ChannelID, &disgord.GetMessagesParams{
				Limit:  100,
				Before: before,
			})
			fmt.Println(messages)
			if err != nil {
				return err
			}
			if len(messages) <= 1 {
				break
			}

			for _, message := range messages {
				channel = fmt.Sprintf(
					"%s%d/%d/%d %d:%d %s : %s\n",
					channel,
					message.Timestamp.Day(),
					message.Timestamp.Month(),
					message.Timestamp.Year(),
					message.Timestamp.Hour(),
					message.Timestamp.Minute(),
					message.Author.Username,
					message.Content,
				)
				before = message.ID
			}
		}

		context.Message.Reply(context.Ctx, bot.Client, "Écriture du fichier")
		_, err = file.Write([]byte(channel))
		context.Message.Reply(context.Ctx, bot.Client, "Done!")
		return
	},
}
