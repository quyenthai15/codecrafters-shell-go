package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input", err)
		os.Exit(1)
	}
	fmt.Printf("%s: command not found", strings.Trim(command, "\n"))
}
