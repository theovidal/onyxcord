package onyxcord

import (
	"io/ioutil"
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
		// Prefix for all the commands (e.g : !)
		Prefix string
		// The color theme of the bot (e.g #b7c1c)
		Color int
	}
	// A list of command categories
	Categories map[string]struct {
		// Their name, displayed in the help command
		Name string
		// Their emoji, displayed in the help command
		Emoji string
	}
	// Some assets user can add
	Assets map[string]string
	// Information to connect to the MongoDB database
	Database struct {
		// Address of the database (e.g : localhost)
		Address string
		// Port of the database (e.g : 1234)
		Port int
		// Username to connect to the database (e.g : onyxcord)
		Username string
		// Password for this username
		Password string
		// Database to use
		Database string
		// Database to connect from (e.g : onyxcord)
		AuthSource string `yaml:"auth_source"`
	}
}

// GetConfig reads a specific bot configuration and parses it into the Config structure
func GetConfig() (config Config, err error) {
	godotenv.Load()

	data, err := OpenFile("./config.yml")
	if err != nil {
		return
	}

	config = Config{}
	err = yaml.Unmarshal(data, &config)

	config.Bot.Token = os.Getenv("DISCORD_TOKEN")
	config.Database.Username = os.Getenv("DATABASE_USERNAME")
	config.Database.Username = os.Getenv("DATABASE_PASSWORD")

	return
}

// OpenFile opens a file from a path
func OpenFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	return
}
