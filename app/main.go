package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	allowedCommands := []string{"exit", "echo", "type"}
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input", err)
			os.Exit(1)
		}
		command = strings.TrimSpace(command)
		action, args, _ := strings.Cut(command, " ")
		if !slices.Contains(allowedCommands, action) {
			fmt.Println(action + ": command not found", )
			continue
		}
		if action == "exit" {
			break
		} else if action == "echo" {
			fmt.Println(args)
			continue
		} else if action == "type" {
			if slices.Contains(allowedCommands, args) {
				fmt.Println(args + " is a shell builtin")
			} else {
				fmt.Println(args + ": not found")
			}
		}
	}
}
