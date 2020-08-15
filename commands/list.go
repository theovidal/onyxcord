package commands

import "github.com/theovidal/onyxcord/lib"

var List = map[string]*lib.Command{
	"help":    &help,
	"poll":    &poll,
	"weather": &weather,
	"ping":    &ping,
}
