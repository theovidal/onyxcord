package commands

import "github.com/BecauseOfProg/boite-a-bois/lib"

var List = map[string]*lib.Command{
	"help":    &help,
	"poll":    &poll,
	"weather": &weather,
	"ping":    &ping,
}
