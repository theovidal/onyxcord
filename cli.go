package onyxcord

import (
	"flag"
	"fmt"
	"log"
)

// ShellCommand is the holder for all CLI commands
type ShellCommand struct {
	// The name of the command
	Name string
	// A short and efficient description for the command
	Description string
	// The flag set to register with the command
	FlagSet *flag.FlagSet
	// The handler function when the command is triggered
	Handler func(*Bot, []string)
}

// String gives a string representation of the command
// Used to get help
func (c *ShellCommand) String() string {
	return fmt.Sprintf("%s - %s", c.Name, c.Description)
}

var DefaultShellCommands = map[string]ShellCommand{
	"help": {
		Name:        "help",
		Description: "Show help",
		FlagSet:     flag.NewFlagSet("help", flag.ExitOnError),
		Handler:     Help,
	},
	"run": {
		Name:        "run",
		Description: "Run the 105chat server",
		FlagSet:     flag.NewFlagSet("run", flag.ExitOnError),
		Handler: func(bot *Bot, strings []string) {
			bot.Run()
		},
	},
	"register": {
		Name:        "register",
		Description: "Register application commands on Discord",
		FlagSet:     flag.NewFlagSet("migrate", flag.ExitOnError),
		Handler:     func(bot *Bot, _ []string) {
			for _, command := range bot.ApplicationCommands {
				if _, err := bot.ApplicationCommandCreate(bot.Config.Bot.AppID, "", command); err != nil {
					log.Panicf("Cannot create %v command: %v", command.Name, err)
				}
			}
			Green.Println("✅ All application commands registered successfully")
		},
	},
}

// Help is a CLI command that gives the command list
func Help(bot *Bot, _ []string) {
	println("───── CLI help ─────\n")

	println("COMMANDS")
	for _, command := range bot.ShellCommands {
		fmt.Println(command.String())
	}
}
