package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/BecauseOfProg/boite-a-bois/lib"

	"github.com/andersfylling/disgord"
	"github.com/exybore/goweather"
)

var Weather = lib.Command{
	Name:        "weather",
	Description: "Obtenir la météo",
	Usage:       "weather <localisation>",
	Category:    "utilities",
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) {
		weatherAPI, err := goweather.NewAPI("33a4c034830b448960d86f6b250bf113", "fr", "metric")
		if err != nil {
			panic(err)
		}

		location := strings.Join(arguments[0:], " ")
		if weather, err := weatherAPI.Current(location); err != nil {
			_, err = context.Message.Reply(
				*bot.Session,
				fmt.Sprint(
					":x: Une erreur inconnue s'est produite : `", err,
					"`\nN'hésitez-pas à contacter le développeur si vous pensez que c'est un bogue !",
				),
			)
			if err != nil {
				log.Fatalf("Error while sending a message : %v", err)
			}
		} else {
			_, _ = context.Message.Respond(
				*bot.Session,
				&disgord.Message{
					Embeds: []*disgord.Embed{
						{
							Title:       fmt.Sprintf("%s :flag_%s:", weather.City.Name, strings.ToLower(weather.City.Country)),
							Description: weather.Conditions.Description,
						},
					},
				},
			)
		}
	},
}
