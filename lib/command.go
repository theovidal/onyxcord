package lib

// Command represents a command that can be executed by a user
type Command struct {
	// Name of the command (e.g : weather)
	Name string
	// Description of the command (e.g : Prints weather for a specific location)
	Description string
	// Usage of the command (e.g : weather <location>)
	Usage string
	// Category of the command, as defined in the configuration (e.g : utilities)
	Category string
	// An alias to another command (e.g : w)
	Alias string
	// Listen for messages from public or private channels.
	// So, the two possible keys inside this array are : private, public
	Listen []string
	// Action to execute if the command is triggered. It's a function with this signature :
	// func(arguments []string, session disgord.Session, context *disgord.MessageCreate)
	Execute interface{}
}
