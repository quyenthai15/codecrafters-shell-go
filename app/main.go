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
	// allowedCommands := []string{"exit", "echo", "type", "pwd"}
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		tokens := strings.Split(command, " ")
		action := tokens[0]
		args := tokens[1:]

		if fn, isBuiltin := BuiltinCmds[action]; isBuiltin {
			fn(args)
			continue
		} else {
			found, _ := FindExecutables(action)
			if found {
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
