package commands

type Timer struct {
	availableCommands []string
}
func setTimer() string{
	return "timer"
}

func NewTimer() *Timer {
	subcmds := []string {"set", "new", "in"}
	return &Timer{subcmds}
}

func (t *Timer) AvailableCommands() []string {
	return t.availableCommands
}

func (t *Timer) FireCommand(as []string) string {
	return "Firing timer command"
}