package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func FindExecutables(command string) (found bool, fullPath string) {
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


func ExecuteCommand(action string, args ...string) {
	out, err := exec.Command(action, args...).Output()
	fmt.Printf("%s", out)
	if err != nil {
		fmt.Println("Command finished with error: ", err)
	}
}
