package commands

import (
	"github.com/ziemerz/gogobot/types"
)

type Command interface {
	SubCommands() map[string] types.Func
}


