package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var BuiltinCmds map[string]func(args []string, w io.Writer)

func init() {
	BuiltinCmds = map[string]func(args []string, w io.Writer){
		"echo": func(args []string, w io.Writer) {
			fmt.Fprintln(w, strings.Join(args, " "))
		},
		"type": func(args []string, w io.Writer) {
			arg := args[0]
			if IsBuiltin(arg) {
				fmt.Fprintln(w, arg+" is a shell builtin")
			} else if absPath, err := exec.LookPath(arg); err == nil {
				fmt.Fprintln(w, arg+" is "+absPath)
			} else {
				fmt.Fprintln(w, arg+": not found")
			}
		},
		"pwd": func(_ []string, w io.Writer) {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(w, "pwd error: ", err)
			} else {
				fmt.Fprintln(w, dir)
			}
		},
		"cd": handleCd,
		"exit": func(_ []string, _ io.Writer) {
			os.Exit(0)
		},
	}
}

func IsBuiltin(cmd string) bool {
	_, ok := BuiltinCmds[cmd]
	return ok
}

func handleCd(args []string, w io.Writer) {
	if len(args) < 1 {
		fmt.Fprintln(w, "cd needs at least 1 argument")
		return
	}
	path := args[0]
	if subPath, found := strings.CutPrefix(path, "~"); found {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		path = homeDir + subPath
	}
	fileInfo, err := os.Stat(path)
	if err != nil || !fileInfo.IsDir() {
		fmt.Fprintf(w, "cd: %s: No such file or directory\n", path)
		return
	}
	if err := os.Chdir(path); err != nil {
		fmt.Fprintf(w, "cd: %s: %v\n", path, err)
	}
}
