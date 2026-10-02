package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)


func readInput() []string {
	fmt.Print("$ ")

	reader := bufio.NewReader(os.Stdin)
	command, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input", err)
		os.Exit(1)
	}
	command = strings.TrimSpace(command)

	args := tokenize(command)
	return args
}


func main() {
	for {
		inputs := readInput()
		if len(inputs) == 0 {
			continue
		}
		action := inputs[0]
		args := inputs[1:]

		args, outFile, err := parseRedirect(args)
		if err != nil {
			fmt.Println("Error parsing redirect: ", err)
			continue
		}
		out := os.Stdout
		if outFile != nil {
			out = outFile
		}

		if cmd, isBuiltin := BuiltinCmds[action]; isBuiltin {
			cmd.Stdout = out
			cmd.Stderr = os.Stderr
			cmd.Run(args)
		} else if _, err := exec.LookPath(action); err == nil {
			cmd := exec.Command(action, args...)
			cmd.Stdout = out
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				var exitErr *exec.ExitError
				// Only print non-exit message
				if isExit := errors.As(err, &exitErr); !isExit {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, action+": command not found")
		}

		// Clean up opened file
		if outFile != nil {
			outFile.Close()
		}
	}
}
