package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BecauseOfProg/boite-a-bois/lib"

	"github.com/andersfylling/disgord"
	"github.com/exybore/goweather"
)

var weather = lib.Command{
	Description: "Obtenir la météo à un endroit précis",
	Usage:       "weather <localisation>",
	Category:    "weather",
	Listen:      []string{"public", "private"},
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		weatherAPI, err := goweather.NewAPI(bot.Config.Assets["weather_api_key"], "fr", "metric")
		if err != nil {
			return
		}

		location := strings.Join(arguments, ",")
		if weather, err := weatherAPI.Current(location); err == nil {
			_ = bot.SendEmbed(
				context.Ctx,
				&disgord.Embed{
					Title: fmt.Sprintf("%s :flag_%s:", weather.City.Name, strings.ToLower(weather.City.Country)),
					Description: fmt.Sprintf("**%s**\n\n"+
						":thermometer: Température : %.1f°C\n"+
						":droplet: Humidité : %.1f%%\n"+
						":cloud: Nuages : %.1f%%\n"+
						":dash: Vent : %.1f km/h",
						strings.Title(weather.Conditions.Description), weather.Conditions.Temperature,
						weather.Conditions.Humidity, weather.Conditions.Clouds,
						weather.Conditions.WindSpeed*3.6),
					Thumbnail: &disgord.EmbedThumbnail{
						URL: weather.Conditions.IconURL,
					},
				},
				context.Message,
			)
		} else {
			return errors.New(":satellite_orbital: Cette localisation est inconnue")
		}
		return
	},
}
