package commands

import "github.com/theovidal/onyxcord/lib"

var List = map[string]*lib.Command{
	"archive":      &archive,
	"help":         &help,
	"ping":         &ping,
	"poll":         &poll,
	"reactionRole": &reactionRole,
	"weather":      &weather,
}
