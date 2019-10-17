package commands

import (
	"log"
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
	Description: "Get the weather",
	Usage:       "weather <localisation>",
	Category:    "utilities",
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) {
		sentenceNumber := rand.Intn(len(pingSentences))
		_, err := context.Message.Reply(*bot.Session, pingSentences[sentenceNumber])
		if err != nil {
			log.Fatalf("Error while sending a message : %v", err)
		}
	},
}
