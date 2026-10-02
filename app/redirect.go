package main

import (
	"fmt"
	"os"
	"slices"
)

func parseRedirect(args []string) (rest []string, w *os.File, err error) {
	redirectIdx := -1
	for i, arg := range slices.Backward(args) {
		if arg == ">" || arg == "1>" {
			redirectIdx = i
			break
		}
	}
	if redirectIdx == -1 {
		return args, nil, nil
	}
	if redirectIdx > -1 {
		outputFile := args[redirectIdx+1]
		rest = args[:redirectIdx]
		w, err = os.Create(outputFile)

		if err != nil {
			fmt.Println("Error creating output file: ", err)
			return nil, nil, err
		}
	}
	return rest, w, nil
}
