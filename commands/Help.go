package commands

import (
	"github.com/ziemerz/gogobot/types"
	"fmt"
)


type Help struct {
	commands map[string]Command
	SubCommands []types.SubCmd
}

func NewHelp() *Help{
	m := GetCmdMap(true).GetCommands(true)
	h := Help{commands: m}
	return &h
}

func (h *Help) AllCommands() string{
	var str string

	for key, elem := range h.commands {
		str += key + ": "
		for _, val := range elem.AvailableCommands() {
			str += "Name: " + val.Name
			str += "Description: " + val.Description
			str += "Help: " + val.Help
		}
	}
	return str
}

func (h *Help) AvailableCommands() []types.SubCmd{
	return h.SubCommands
}

func specificCommand(cmd string) string{
	return cmd + " specific help"
}

func (h *Help) FireCommand(command []string) string {
	fmt.Println(command)
	if len(command) == 1 {
		return h.AllCommands()
	}
	return specificCommand(command[1])
}