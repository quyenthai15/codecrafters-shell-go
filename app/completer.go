package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/chzyer/readline"
)


type BuiltinCompleter struct {
	completer readline.AutoCompleter
}


func NewCustomCompleter() *BuiltinCompleter {
	var items []readline.PrefixCompleterInterface

	// List builtin commands
	for key := range BuiltinCmds {
		items = append(items, readline.PcItem(key))
	}

	// List all executables from PATH
	pathEnv := os.Getenv("PATH")
	for _, dir := range filepath.SplitList(pathEnv) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}
		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			if !info.IsDir() && info.Mode().Perm()&0111 != 0 {
				items = append(items, readline.PcItem(entry.Name()))
			}
		}
	}

	items = slices.Concat(items)
	return &BuiltinCompleter{
		completer: readline.NewPrefixCompleter(items...),
	}
}

func (c *BuiltinCompleter) Do(line []rune, pos int) (newLine [][]rune, offset int) {
	newLine, offset = c.completer.Do(line, pos)
	if offset == 0 {
		fmt.Print("\x07")
		return nil, 0
	}

	return newLine, offset
}
