package commands

import (
	"github.com/theovidal/onyxcord/lib"
)

var List = map[string]*lib.Command{
	"ping":    &ping,
	"poll":    &poll,
	"weather": &weather,
}
