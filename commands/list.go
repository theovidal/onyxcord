package commands

import "github.com/BecauseOfProg/boite-a-bois/lib"

var List = map[string]*lib.Command{
	"weather": &weather,
	"ping":    &ping,
}
