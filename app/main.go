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
		if command == "exit" {
			break
		} else if phrase, found := strings.CutPrefix(command, "echo "); found {
			fmt.Println(phrase)
			continue
		}
		fmt.Println(command + ": command not found", )
	}
}
