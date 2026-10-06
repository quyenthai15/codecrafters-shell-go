package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)


func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:       "$ ",
		AutoComplete: completer,
		HistoryFile:  "/tmp/readline.tmp",
		HistoryLimit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()
	rl.CaptureExitSignal()

	for {
		line, err := rl.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		inputs := tokenize(strings.TrimSpace(line))

		if len(inputs) == 0 {
			continue
		}
		action := inputs[0]
		args := inputs[1:]

		args, outFile, errFile, err := parseRedirect(args)
		if err != nil {
			fmt.Println("Error parsing redirect: ", err)
			continue
		}
		stdout := os.Stdout
		stderr := os.Stderr
		if outFile != nil {
			stdout = outFile
		}
		if errFile != nil {
			stderr = errFile
		}

		if cmd, isBuiltin := BuiltinCmds[action]; isBuiltin {
			cmd.Stdout = stdout
			cmd.Stderr = stderr
			cmd.Run(args)
		} else if _, err := exec.LookPath(action); err == nil {
			cmd := exec.Command(action, args...)
			cmd.Stdout = stdout
			cmd.Stderr = stderr
			if err := cmd.Run(); err != nil {
				var exitErr *exec.ExitError
				// Only print non-exit message
				if isExit := errors.As(err, &exitErr); !isExit {
					fmt.Fprintln(stderr, err)
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, action+": command not found")
		}

		// Clean up opened file
		if outFile != nil {
			outFile.Close()
		}
		if errFile != nil {
			errFile.Close()
		}
	}
}
