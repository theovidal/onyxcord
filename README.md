<div align="center">
    <img src="./onyxcord.png" alt="onyxcord"/>
    <h1>onyxcord</h1>
    <h3>Discord bots with Go as straightforward as possible.</h3>
    <a href="https://pkg.go.dev/github.com/theovidal/onyxcord">Documentation</a> — <a href="https://discord.gg/QGGSTXy">Discord server</a> — <a href="./LICENSE">License</a>
</div>

**⚠ This is only a prototype and for a personal use. Don't plan to make huge bots with this library.**

## 🔧 Setup

Get the dependency from the source:

```bash
go get -u github.com/theovidal/onyxcord
```

In your code, create the bot instance:

```go
bot := onyxcord.RegisterBot("MyBot", true)
```

You can then register commands:

```go
pingCommand := *onyxcord.Command{
	// ...
}
bot.RegisterCommand("ping", &pingCommand)
```

Specify the intents of your bot, so it can receive proper events:

```go
bot.Client.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages)
```

Finally, connect it to Discord with the `Run` method! See more on the [documentation](https://pkg.go.dev/github.com/theovidal/onyxcord) of the library.

## 💻 Development

TODO

## 📜 Credits

- Library: [discordgo](https://github.com/bwmarrin/discordgo)
- Maintainer: [Théo Vidal](https://github.com/theovidal)

## 🔐 License

[GNU GPL v3](./LICENSE)
