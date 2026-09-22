package main

import (
	"fmt"
	"os"
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
