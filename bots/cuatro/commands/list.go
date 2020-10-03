package commands

import (
	"github.com/theovidal/onyxcord/lib"
)

var List = map[string]*lib.Command{
	"archive":      &archive,
	"reactionRole": &reactionRole,
}
