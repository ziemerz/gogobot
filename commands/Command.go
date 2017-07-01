package commands

import "github.com/ziemerz/gogobot/types"

type Command interface {
	AvailableCommands() []types.SubCmd
	FireCommand([]string) string
}


