package main

import "github.com/ziemerz/gogobot/commands"

type CommandHandler struct {
	commandsMap map[string] commands.Command
}

func NewCommandHandler() CommandHandler{
	commandsMap := make(map[string] commands.Command)
	commandsMap["random"] = commands.NewRandom()
	// commandsMap["timer"] = commands.Timer{}
	return CommandHandler{commandsMap}
}

func (cmdHandler *CommandHandler) HandleCommand(command string, subcommand string) string{
	if command == "random" {
		if subcommand == "cat" {
			return cmdHandler.commandsMap["random"].SubCommands()["cat"]()
		} else if subcommand == "gif" {
			return cmdHandler.commandsMap["random"].SubCommands()["gif"]()
		} else if subcommand == "dog" {
			return cmdHandler.commandsMap["random"].SubCommands()["dog"]()
		}
	}
	return "Command not found"
}

func (cmdHandler *CommandHandler) HandleCommandNew(command []string) string {
	if command[0] == "cat" {
		return cmdHandler.commandsMap["random"].SubCommands()["cat"]()
	} else if command[0] == "dog" {
		return cmdHandler.commandsMap["random"].SubCommands()["dog"]()
	} else if command[0] == "timer" {
		return cmdHandler.commandsMap["timer"].SubCommands()[command[1]]()
	}
	return "HI!"
}
