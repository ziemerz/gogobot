package main

import (
	"github.com/ziemerz/gogobot/commands"
	"github.com/ziemerz/gogobot/types"
)

type CommandHandler struct {
	commandsMap map[string] commands.Command
}

func NewCommandHandler() CommandHandler{
	cmdMap := types.GetCmdMap()
	return CommandHandler{cmdMap.GetCommands()}
}

func (cmdHandler *CommandHandler) HandleCommand(command []string) string {
	if len(command) >= 1 {
		if command[0] == "random" {
			return cmdHandler.commandsMap["random"].FireCommand(command)
		} else if command[0] == "timer" {
			return cmdHandler.commandsMap["timer"].FireCommand(command)
		}
	}
	return "Command not found"
}
