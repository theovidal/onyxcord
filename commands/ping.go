package commands

import (
	"math/rand"

	"github.com/BecauseOfProg/boite-a-bois/lib"
	"github.com/andersfylling/disgord"
)

var pingSentences = []string{
	"Pong ! :smirk:",
	"BONJOUR",
	"gildas il sait trop bien grimper putain, exybore est jaloux de ouf",
	"t'façon tout le monde sait que Java ça pue",
	"https://github.com/exybore :wink:",
	`On dit "iPhone dix" et non "iPhone ixe" :rage:`,
}

var ping = lib.Command{
	Description: "Tester si le robot répond correctement",
	Usage:       "ping",
	Category:    "utilities",
	Show:        false,
	Listen:      []string{"public", "private"},
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		sentenceNumber := rand.Intn(len(pingSentences))
		_, err = context.Message.Reply(context.Ctx, *bot.Session, pingSentences[sentenceNumber])
		return
	},
}
