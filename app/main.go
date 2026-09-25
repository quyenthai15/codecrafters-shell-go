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


func tokenize(text string) (tokens []string) {
	text = strings.TrimSpace(text)
	isSpaceStarted := false
	isSingleQuoteStarted := false
	for idx, char := range text {
		if len(tokens) == 0 {
			tokens = append(tokens, "")
		}
		// Checking single quoted
		if char == '\'' {
			isSingleQuoteStarted = !isSingleQuoteStarted
			isSpaceStarted = false
			continue
		}
		if isSingleQuoteStarted {
			tokens[len(tokens) - 1] += string(char)
			continue
		}
		// treating space
		if char == ' ' {
			// ignoreing consecutive spaces
			if idx >= 1 && text[idx - 1] == ' ' {
				continue
			}

			if !isSpaceStarted {
				isSpaceStarted = true
				tokens = append(tokens, "")
			}
		} else {
			isSpaceStarted = false
			tokens[len(tokens) - 1] += string(char)
		}
	}
	return tokens
}
