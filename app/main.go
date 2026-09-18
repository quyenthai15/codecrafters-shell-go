package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		if action == "exit" {
			break
		} else if action == "echo" {
			fmt.Println(args)
			continue
		} else if action == "type" {
			if slices.Contains(allowedCommands, args) {
				fmt.Println(args + " is a shell builtin")
			} else if found, fullPath := findExecutables(args); found {
				fmt.Println(args + " is " + fullPath)
			} else {
				fmt.Println(args + ": not found")
			}
		} else {
			found, _ := findExecutables(action)
			if found {
				arguments := []string{}
				if len(args) > 0 {
					arguments = strings.Split(args, " ")
				}
				out, err := exec.Command(action, arguments...).Output()
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

func findExecutables(command string) (found bool, fullPath string) {
	path := os.Getenv("PATH")

	for dir := range strings.SplitSeq(path, ":") {
		entries, error := os.ReadDir(dir)
		if error != nil || len(entries) == 0 {
			continue
		}
		for _, e := range entries {
			info, error := e.Info()
			if error == nil && info.Name() == command && strings.ContainsRune(info.Mode().String(), 'x'){
				found = true
				fullPath = fmt.Sprintf("%s/%s", dir, info.Name())
				return found, fullPath
			}
		}
	}
	return found, fullPath
}
