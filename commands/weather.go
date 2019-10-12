package commands

import (
	"fmt"

	"github.com/BecauseOfProg/boite-a-bois/lib"

	"github.com/exybore/goweather"
	"github.com/andersfylling/disgord"
)

var Weather = lib.Command{
	Name: "weather",
	Description: "Obtenir la météo",
	Usage: "weather <localisation>",
	Category: "utilities",
	Execute: func(arguments []string, session disgord.Session, context *disgord.MessageCreate) {
		weatherAPI, err := goweather.NewAPI("33a4c034830b448960d86f6b250bf113", "fr", "metric")
		if err != nil {
			panic(err)
		}

		if weather, err := weatherAPI.Current(arguments[0]); err != nil {
			context.Message.Reply(
				session,
				fmt.Sprint(
					":x: Une erreur inconnue s'est produire :", err,
					"\nN'hésitez-pas à contacter le développeur si vous pensez que c'est un bug !",
				),
			)
		} else {
			fmt.Println(weather)
			context.Message.Reply(session, fmt.Sprint(weather.City.Name, "\n", weather.Conditions.Description))
		}
	},
}