package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		if command == "" {
			continue
		}
		tokens := strings.Split(command, " ")
		action := tokens[0]
		args := tokens[1:]

		if fn, isBuiltin := BuiltinCmds[action]; isBuiltin {
			fn(args)
			continue
		} else {
			found, _ := FindExecutables(action)
			if found {
				ExecuteCommand(action, args...)
			} else {
				fmt.Println(action + ": command not found", )
			}
		}
	}
}
