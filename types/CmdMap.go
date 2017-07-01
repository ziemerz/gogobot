package types

import (
	"github.com/ziemerz/gogobot/commands"
	"sync"
)

type CmdMap struct {
	cmds map[string] commands.Command
}

var instance *CmdMap
var once sync.Once

func GetCmdMap() *CmdMap{
	once.Do(func() {
		mh := make(map[string] commands.Command)
		mh["random"] = commands.NewRandom()
		mh["timer"] = commands.NewTimer()
		instance = &CmdMap{mh}
	})
	return instance
}

func (cm *CmdMap) GetCommand(cmd string) commands.Command {
	return instance.cmds[cmd]
}

func (cm *CmdMap) GetCommands() map[string] commands.Command {
	return instance.cmds
}
