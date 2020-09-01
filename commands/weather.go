package commands

import (
	"errors"
	"fmt"
	"strings"

	"github.com/theovidal/onyxcord/lib"

	"github.com/andersfylling/disgord"
	"github.com/exybore/goweather"
)

var weather = lib.Command{
	Description:    "Obtenir la météo à un endroit précis",
	Usage:          "weather <localisation>",
	Category:       "weather",
	ListenInPublic: true,
	ListenInDM:     true,
	Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) (err error) {
		weatherAPI, err := goweather.NewAPI(bot.Config.Assets["weather_api_key"], "fr", "metric")
		if err != nil {
			return
		}

		location := strings.Join(arguments, ",")
		weather, err := weatherAPI.Current(location)
		if err != nil {
			return errors.New(":satellite_orbital: Cette localisation est inconnue")
		}

		err = bot.SendEmbed(
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
		return
	},
}
