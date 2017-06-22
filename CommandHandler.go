package main

import "github.com/ziemerz/gogobot/commands"

type CommandHandler struct {
	commandsMap map[string] commands.Command
}

func NewCommandHandler() CommandHandler{
	commandsMap := make(map[string] commands.Command)
	commandsMap["random"] = commands.NewRandom()
	return CommandHandler{commandsMap}
}

func (cmdHandler *CommandHandler) HandleCommand(command string, subcommand string) string{
	if command == "random" {
		if subcommand == "cat" {
			return cmdHandler.commandsMap["random"].SubCommands()["cat"]()
		} else if subcommand == "gif" {
			return cmdHandler.commandsMap["random"].SubCommands()["gif"]()
		}
	}
	return "Command not found"
}
