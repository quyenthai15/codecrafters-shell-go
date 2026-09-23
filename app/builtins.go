package main

import (
	"fmt"
	"os"
	"strings"
)

var BuiltinCmds map[string]func(args []string)

func init() {
	BuiltinCmds = map[string]func(args []string){
		"echo": func(args []string) {
			fmt.Println(strings.Join(args, " "))
		},
		"type": func(args []string) {
			arg := args[0]
			if IsBuiltin(arg) {
				fmt.Println(arg + " is a shell builtin")
			} else if found, fullPath := FindExecutables(arg); found {
				fmt.Println(arg + " is " + fullPath)
			} else {
				fmt.Println(arg + ": not found")
			}
		},
		"pwd": func(_ []string) {
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("pwd error: ", err)
			} else {
				fmt.Println(dir)
			}
		},
		"cd": func(args []string) {
			if len(args) < 1 {
				fmt.Println("cd needs at least 1 argument")
				return
			}
			path := args[0]
			if subPath, found := strings.CutPrefix(path, "~"); found {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println(err)
					return
				}
				path = homeDir + subPath
			}
			fileInfo, err := os.Stat(path)
			if err != nil || !fileInfo.IsDir() {
				fmt.Printf("cd: %s: No such file or directory\n", path)
				return
			}
			if err := os.Chdir(path); err != nil {
				fmt.Printf("cd: %s: %v\n", path, err)
			}
		},
		"exit": func(_ []string) {
			os.Exit(0)
		},
	}
}

func IsBuiltin(cmd string) bool {
	_, ok := BuiltinCmds[cmd]
	return ok
}
