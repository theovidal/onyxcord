package onyxcord

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

// Config represents the configuration stored in the file
type Config struct {
	// Information for the development
	Dev struct {
		// Activates the debug mode, so the bot logs everything
		Debug bool
		// Version of the bot
		Version string
		// Information about the maintainer
		Maintainer struct {
			Name string
			Link string
		}
	}
	// Information about the bot itself
	Bot struct {
		// Its name
		Name string
		// The token to connect to Discord
		Token string
		// Discord's client ID
		ID string
		// An illustration, for example its logo
		Illustration string
		// A website on which users can go for further information
		Website string
		// The link to invite the bot on a Discord server
		InviteLink string `yaml:"invite_link"`
		// Link to the license of the bot (if there's one)
		License string
		// Link to support the development (if there's one)
		Donate string
		// The color theme of the bot (e.g: 7976509)
		Color int
		// The color associated to an error (e.g: 12000284)
		ErrorColor int `yaml:"error_color"`
	}
	// Some assets user can add
	Assets map[string]string
	// Information to connect to the MongoDB database
	Database struct {
		// Address of the database (e.g: localhost)
		Address string
		// Port of the database (e.g: 1234)
		Port int
		// Username to connect to the database (e.g: onyxcord)
		Username string
		// Password for this username
		Password string
		// Database to use
		Database string
		// Database to connect from (e.g: onyxcord)
		AuthSource string `yaml:"auth_source"`
	}
	Cache struct {
		// Address of the cache (e.g: localhost)
		Address string
		// Port of the cache (e.g: 6379)
		Port string
		// Password to access the cache (optional but strongly recommended)
		Password string
	}
}

// GetConfig reads a specific bot configuration and parses it into the Config structure
func GetConfig() (config Config, err error) {
	godotenv.Load()

	var path string
	if os.Getenv("ONYXCORD_ENV") == "development" {
		path = "./config.dev.yml"
	} else {
		path = "./config.yml"
	}

	data, err := OpenFile(path)
	if err != nil {
		return
	}

	config = Config{}
	err = yaml.Unmarshal(data, &config)

	config.Bot.Token = os.Getenv("DISCORD_TOKEN")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Cache.Password = os.Getenv("CACHE_PASSWORD")

	return
}
