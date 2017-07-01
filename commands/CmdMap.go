package commands

import (
	"sync"
)

type cmdMap struct {
	cmds map[string] Command
}

var instance *cmdMap
var onceMap sync.Once
var onceHelp sync.Once
var noHelpInstance *cmdMap

func GetCmdMap(isHelp bool) *cmdMap{
	if !isHelp {
		onceMap.Do(func() {
			mh := make(map[string]Command)
			mh["random"] = NewRandom()
			mh["timer"] = NewTimer()
			mh["help"] = NewHelp()
			instance = &cmdMap{mh}
		})
		return instance
	} else {
		onceHelp.Do(func() {
			mh := make(map[string]Command)
			mh["random"] = NewRandom()
			mh["timer"] = NewTimer()
			noHelpInstance = &cmdMap{mh}
		})
		return noHelpInstance
	}
}

func (cm *cmdMap) GetCommand(cmd string) Command {
	return instance.cmds[cmd]
}

func (cm *cmdMap) GetCommands(isHelp bool) map[string] Command {
	if !isHelp {
		return instance.cmds
	} else {
		return noHelpInstance.cmds
	}
}

