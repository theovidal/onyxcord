package commands

import (
	"errors"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/theovidal/onyxcord/lib"
	"strings"
)

var pollChoices = map[string][]string{
	"shapes":         strings.Split("🔴 🟤 🟠 🟣 🟡 🔵 🟢 ⚫ ⚪ 🟥 🟫 🟧 🟪 🟨 🟦 🟩 ⬛ ⬜ 🔶 🔺", " "),
	"numbers":        strings.Split("0️⃣ 1️⃣ 2️⃣ 3️⃣ 4️⃣ 5️⃣ 6️⃣ 7️⃣ 8️⃣ 9️⃣ 🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮", " "),
	"letters":        strings.Split("🇦 🇧 🇨 🇩 🇪 🇫 🇬 🇭 🇮 🇯 🇰 🇱 🇲 🇳 🇴 🇵 🇶 🇷 🇸 🇹", " "),
	"food":           strings.Split("🍎 🍍 🍇 🥐 🥗 🥪 🍕 🥓 🍜 🥘 🍧 🍩 🍰 🍬 🍭 ☕ 🧃 🍵 🍾 🍸", " "),
	"faces":          strings.Split("😄 😋 😎 😂 🥰 😎 🤔 🙄 😑 🤨 😮 😴 😛 😤 🤑 😭 😨 🥵 🥶 😷", " "),
	"animals":        strings.Split("🐔 🐴 🐸 🐷 🐗 🐰 🐹 🦊 🐶 🐼 🦓 🐁 🐘 🐢 🐍 🐳 🦐 🐠 🦢 🦜", " "),
	"transportation": strings.Split("🚗 🚓 🚌 🚚 🚜 🚅 🚋 🚇 🚠 ✈ 🚁 🚀 🚢 🛹 🚲 🛴 🛵 🚑 🚒 🦽", " "),
}

var poll = lib.Command{
	Description: "Organiser un vote",
	Usage:       "poll <template>,<question>,[choix...]",
	Category:    "utilities",
	Listen:      []string{"public"},
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		err = bot.Client.DeleteMessage(context.Ctx, context.Message.ChannelID, context.Message.ID)
		if err != nil {
			return
		}

		if len(arguments[2:]) > 20 || len(arguments[2:]) == 0 {
			return errors.New("Le nombre de réponses doit être compris entre 1 et 22")
		}

		var template []string
		if arguments[1] == "" {
			template = pollChoices["letters"]
		} else {
			_, ok := pollChoices[arguments[1]]
			if !ok {
				return errors.New(
					"Le modèle de choix est invalide. " +
						"Les choix possibles sont : `shapes`, `numbers`, `letters`, `food`, `faces`, `transportation`",
				)
			} else {
				template = pollChoices[arguments[1]]
			}
		}

		var choices string
		for index, value := range arguments[2:] {
			choices += fmt.Sprintf("%s %s\n", template[index], value)
		}

		userAvatar, _ := context.Message.Author.AvatarURL(64, false)
		message := disgord.Embed{
			Title:       fmt.Sprintf("**📊 Sondage :** %s", arguments[0]),
			Description: choices,
			Author: &disgord.EmbedAuthor{
				Name:    context.Message.Author.Username,
				IconURL: userAvatar,
			},
		}

		sentMessage, err := bot.Client.SendMsg(
			context.Ctx,
			context.Message.ChannelID,
			&disgord.CreateMessageParams{
				Embed: lib.MakeEmbed(
					bot.Config,
					&message,
				),
			},
		)
		if err != nil {
			return
		}

		for index, _ := range arguments[2:] {
			err = sentMessage.React(
				context.Ctx,
				*bot.Session,
				template[index],
			)
			if err != nil {
				return
			}
		}

		return
	},
}
