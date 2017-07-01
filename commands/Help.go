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
		str += "```" +key + ": \n"
		for _, val := range elem.AvailableCommands() {
			str += "\tSub command: " + val.Name + "\n"
			str += "\t\tDescription: " + val.Description + "\n"
			str += "\t\tExample: " + val.Example + "\n"
			str += "------------\n"
		}
		str += "```"
	}
	return str
}

func (h *Help) AvailableCommands() []types.SubCmd{
	return h.SubCommands
}

func (h *Help) specificCommand(cmd string) string{
	var str string = "Help for " + cmd + "\n ```"
	for _, command := range h.commands[cmd].AvailableCommands() {
		str += "\tSub command: " + command.Name + "\n"
		str += "\t\tDescription: " + command.Description + "\n"
		str += "\t\tExample: " + command.Example + "\n"
	}
	return str + "```"
}

func (h *Help) FireCommand(command []string) string {
	fmt.Println(command)
	if len(command) == 1 {
		return h.AllCommands()
	}
	return h.specificCommand(command[1])
}