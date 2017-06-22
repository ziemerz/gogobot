package commands

import (
	"github.com/ziemerz/gogobot/types"
)

type Random struct {
	subCommands map[string]types.Func
}

func randomCat() string {
	return "Random cat"
}

func randomGif() string {
	return "Random gif"
}

func (r *Random) SubCommands() map[string]types.Func {
	return r.subCommands
}

func NewRandom() *Random{
	subcmds := make(map[string]types.Func)
	subcmds["cat"] = randomCat
	subcmds["gif"] = randomGif
	return &Random{subcmds}
}

