package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

const bellChar = '\x07'

var completer = readline.NewPrefixCompleter(
	readline.PcItem("echo"),
	readline.PcItem("exit"),
	readline.PcItem("pwd"),
	readline.PcItem("cd"),
	readline.PcItem("type"),

	readline.PcItemDynamic(func(s string) []string {
			fmt.Print(string(bellChar))
			return []string{}
		},
	),
)
