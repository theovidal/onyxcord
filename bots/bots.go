package bots

import (
	"github.com/theovidal/onyxcord/bots/boite_a_bois"
	"github.com/theovidal/onyxcord/bots/cuatro"
	"github.com/theovidal/onyxcord/lib"
)

var Bots = map[*lib.Bot]bool{
	boite_a_bois.Install(): true,
	cuatro.Install():       true,
}
