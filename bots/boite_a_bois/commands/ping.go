package commands

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"

	"github.com/theovidal/onyxcord/lib"
)

var pingSentences = []string{
	"Pong ! :smirk:",
	"BONJOUR",
	"gildas il sait trop bien grimper putain, théo est jaloux de ouf",
	"t'façon tout le monde sait que Java ça pue",
	"https://github.com/theovidal Meilleur profil GitHub du monde :wink:",
	`On dit "iPhone dix" et non "iPhone ixe" :rage:`,
	"*prend la voix de kernoeb* Heyyy",
	"jaaj",
	"Mais j'ai l'impression de rendre l'aléatoire encore plus aléatoire",
	"OK boomer",
}

func Ping() *lib.Command {
	return &lib.Command{
		Description: "Tester si le robot répond correctement",
		Usage:       "ping",
		Category:    "utilities",
		Show:        false,
		ListenInDM:  true,
		Execute: func(arguments []string, bot lib.Bot, message *discordgo.MessageCreate) (err error) {
			sentenceNumber := rand.Intn(len(pingSentences))
			_, err = bot.Client.ChannelMessageSend(message.ChannelID, pingSentences[sentenceNumber])
			return
		},
	}
}
