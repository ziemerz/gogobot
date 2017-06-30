package commands

type Command interface {
	AvailableCommands() []string
	FireCommand([]string) string
}


