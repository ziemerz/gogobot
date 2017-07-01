package main

import (
	"github.com/ziemerz/gogobot/commands"
)

type CommandHandler struct {
	commandsMap map[string] commands.Command
}

func NewCommandHandler() CommandHandler{
	cmdMap := commands.GetCmdMap(false)
	return CommandHandler{commandsMap:cmdMap.GetCommands(false)}
}

func (cmdHandler *CommandHandler) HandleCommand(command []string) string {
	if len(command) >= 1 {
		switch command[0] {
		case "random":
			return cmdHandler.commandsMap["random"].FireCommand(command)
		case "timer":
			return cmdHandler.commandsMap["timer"].FireCommand(command)
		case "help":
			return cmdHandler.commandsMap["help"].FireCommand(command)
		}
	}
	return "Command not found"
}
