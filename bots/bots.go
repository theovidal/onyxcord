package bots

import (
	"github.com/theovidal/onyxcord/bots/boite_a_bois"
	"github.com/theovidal/onyxcord/bots/carrbridge"
	"github.com/theovidal/onyxcord/bots/cuatro"
	"github.com/theovidal/onyxcord/lib"
)

var Bots = map[string]lib.Bot{
	"boite_a_bois": boite_a_bois.Install(),
	"carrbridge":   carrbridge.Install(),
	"cuatro":       cuatro.Install(),
}
