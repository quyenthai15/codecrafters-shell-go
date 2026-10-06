package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

var allCmds []readline.PrefixCompleterInterface

var completer = readline.NewPrefixCompleter(
	readline.PcItem("echo"),
	readline.PcItem("exit"),
	readline.PcItem("pwd"),
	readline.PcItem("cd"),
	readline.PcItem("type"),
)

var rlConfig = readline.Config{
	Prompt:       "$ ",
	AutoComplete: &BuiltinCompleter{},
	HistoryFile:  "/tmp/readline.tmp",
	HistoryLimit: 20,
}

type BuiltinCompleter struct {}

func (c *BuiltinCompleter) Do(line []rune, pos int) (newLine [][]rune, offset int) {
	newLine, offset = completer.Do(line, pos)
	if offset == 0 {
		fmt.Print(string('\x07'))
		return nil, 0
	}

	return newLine, offset
}

func init() {
	for key := range BuiltinCmds {
		allCmds = append(allCmds, readline.PcItem(key))
	}
}
