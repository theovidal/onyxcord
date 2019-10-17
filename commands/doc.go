// Package commands contains all the commands that the bot supports.
// Each one is a variable of type lib.Command :
//   var weather = lib.Command{
//	   Description: "Get the weather",
//	   Usage:       "weather <localisation>",
//	   Category:    "utilities",
//	   Execute: func(arguments []string, bot lib.Bot, context *disgord.MessageCreate) {
//       // action to execute when the command is triggered
//     },
//   }
// Then, each command is inside the commands.List variable :
//   var List = map[string]*lib.Command{
//	   "weather": &Weather,
//   }
package commands
