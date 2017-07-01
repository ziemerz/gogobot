package commands

import (
	"github.com/ziemerz/gogobot/utility"
	"github.com/ziemerz/gogobot/types"
)


type TimerWrapper struct {
	Timer Timer	`json:"timer"`
}

type Timer struct {
	SubCommands []types.SubCmd `json:"subcommands"`
}

func setTimer() string{
	return "timer"
}

func NewTimer() *Timer {
	var tm TimerWrapper
	utility.GetJsonFromFile("commands.json", &tm)
	return &tm.Timer
}

func (t *Timer) AvailableCommands() []types.SubCmd {
	return t.SubCommands
}

func (t *Timer) FireCommand(as []string) string {
	return "Sup timer?"
}