package main

import (
	"fmt"
	"os"
)

func parseRedirect(args []string) (rest []string, stdoutFile *os.File, stderrFile *os.File, err error) {
	var stdout, stderr string
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case ">", "1>":
			if len(args) > i+1 {
				stdout = args[i+1]
				i++
			}
		case "2>":
			if len(args) > i+1 {
				stderr = args[i+1]
				i++
			}
		default:
			rest = append(rest, args[i])
		}
	}

	if stdout == "" && stderr == "" {
		return args, nil, nil, nil
	}

	if stdout != "" {
		stdoutFile, err = os.Create(stdout)
		if err != nil {
			fmt.Println("Error creating output file: ", err)
			return nil, nil, nil, err
		}
	}

	if stderr != "" {
		stderrFile, err = os.Create(stderr)
		if err != nil {
			fmt.Println("Error creating err file: ", err)
			return nil, nil, nil, err
		}
	}
	return rest, stdoutFile, stderrFile, nil
}
