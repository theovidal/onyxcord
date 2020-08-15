<div align="center">
  <img src="https://cdn.becauseofprog.fr/v2/projects/boite-a-bois.png" width="200" alt="logo">
  <h1>Boite à bois</h1>
  <h6><i>literally "wood box"</i></h6>
  <h3>The French Discord bot of the BecauseOfProg.</h3>
  <a href="https://becauseofprog.fr">Website</a> - <a href="https://discord.becauseofprog.fr">Discord server</a>
</div>

- [🌈 Features](#-features)
- [📲 Requirements](#-requirements)
- [⏩ First start](#-first-start)
- [🔧 Creating commands](#-creating-commands)
- [📚 Creating database models](#-creating-database-models)
- [📜 Credits](#-credits)
- [🔐 License](#-license)

## 🌈 Features

- Play games many games, singleplayer or multiplayer
- Get the current weather and weather forecast for any city in the world
- Send very nice GIF
- Search for posts and users of [our blog](https://becauseofprog.fr)

## 📲 Requirements

- Go 1.12+
- A MongoDB database
- Dependencies listed in [go.mod](go.mod)

## ⏩ First start

To use it, you must setup config.yml. You can help you with the config.sample.yml file, which contains a sample configuration.
When it's configured, you can start the bot with this command :

```ruby
go run main.rb
```

After that, the bot is ready and you can add it to your guild ([Guide](https://discordapp.com/developers/docs/topics/oauth2#bot-authorization-flow))

## 🔧 Creating commands

To create a new command, you first must create a file under the commands directory of the name of your command.
In this file, you have to set the package to commands and import the lib package of this repository.
Create a variable of type `lib.Command` and fill in the fields.
Here is an example :

```go
package commands

import (
  "github.com/andersfylling/disgord"
  "github.com/theovidal/onyxcord/lib"
)

var weather = lib.Command{
  Description: "Get the weather",
  Usage:       "weather <localisation>",
  Category:    "utilities",
  Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) {
    // action to execute when the command is triggered
  },
}
```

Finally, add your freshly created command to the list in the `commands/list.go` file :

```go
var List = map[string]*lib.Command{
  "weather": &weather,
  "ping":    &ping,
}
```

## 📚 Creating database models

**Todo : update to the Go bot**

All the models are defined in two files :

- `db/models` : the Go structures
- `db/schemas` : the database schema, if required

The schemas are in the YML format. All is documented [on the mongocore gem page](https://github.com/fugroup/mongocore)

## 📜 Credits

- Library : [DiscordRB](https://github.com/meew0/discordrb)
- Developers :
  - [Whaxion](https://github.com/whaxion) : old Ruby bot
  - [Exybore](https://github.com/exybore) : maintainor, actual developer

## 🔐 License

[GNU GPL v3](LICENSE)
