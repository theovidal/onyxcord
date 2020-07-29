package lib

import (
	"fmt"
	"strings"
)

// Options for command listening
var ListenTypes = map[string]string{
	"public":  "salon public",
	"private": "message privé",
}

// Command represents a command that can be executed by a user
type Command struct {
	// Description of the command (e.g : Prints weather for a specific location)
	Description string
	// Usage of the command (e.g : weather <location>)
	Usage string
	// Category of the command, as defined in the configuration (e.g : utilities)
	Category string
	// An alias to another command (e.g : w)
	Alias string
	// Choose if the command is shown in the help or not
	Show bool
	// Listen for messages from public or private channels.
	// So, the two possible keys inside this array are : private, public
	Listen []string
	// Lock the command only for certain channels
	Channels []int
	// Lock the command only for certain user roles
	Roles []int
	// Lock the command only for certain members on the server
	Members []int
	// Action to execute if the command is triggered. It's a function with this signature :
	// func(arguments []string, session disgord.Session, context *disgord.MessageCreate)
	Execute interface{}
}

// Prettify returns a string with information about a command, ready to be printed to the user
func (command Command) Prettify(name string, prefix string) (prettified string) {
	prettified = fmt.Sprintf("● **%s%s** : %s\nUtilisation : `%s`", prefix, name, command.Description, command.Usage)
	if command.Alias != "" {
		prettified += fmt.Sprintf("\nAlias : `%s`", command.Alias)
	}
	prettified += fmt.Sprintf("\nS'exécute en")
	for _, listenType := range command.Listen {
		prettified += fmt.Sprintf(" %s et", ListenTypes[listenType])
	}
	prettified = strings.TrimSuffix(prettified, " et")
	return
}
