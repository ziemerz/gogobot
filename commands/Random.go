package commands

type Func func() string
type Random struct {
	Subcommands map[string]Func
}

func randomCat() string {
	return "Random cat"
}

func NewRandom() *Random{
	subcommands := map[string]Func{
		"cat": randomCat,
	}

	return &Random{subcommands}
}

