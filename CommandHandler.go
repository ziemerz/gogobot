package main

import "github.com/ziemerz/gogobot/commands"

type CommandHandler struct {
	commandsMap map[string] commands.Command
}

func NewCommandHandler() CommandHandler{
	cmds := make(map[string] commands.Command)
	cmds["random"] = commands.NewRandom()
	cmds["timer"] = commands.NewTimer()
	return CommandHandler{cmds}
}

func (cmdHandler *CommandHandler) HandleCommand(command []string) string {
	if len(command) >= 2 {
		if command[0] == "random" {
			return cmdHandler.commandsMap["random"].FireCommand(command[1:])
		} else if command[0] == "timer" {
			return cmdHandler.commandsMap["timer"].FireCommand(command[1:])
		}
	}
	return "Command not found"
}
