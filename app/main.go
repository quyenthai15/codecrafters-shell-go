package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		tokens := strings.SplitN(command, " ", 2)
		action := tokens[0]
		var args []string
		if len(tokens) > 1 {
			args = tokenize(tokens[1])
		}

		if handler, isBuiltin := BuiltinCmds[action]; isBuiltin {
			handler(args)
			continue
		} else {
			if _, err := exec.LookPath(action); err == nil {
				out, err := exec.Command(action, args...).Output()
				fmt.Printf("%s", out)
				if err != nil {
					fmt.Println("Command finished with error: ", err)
				}
			} else {
				fmt.Println(action + ": command not found", )
			}
		}
	}
}
