package main

import (
	"fmt"
	"strings"
)


type lexState int

const (
	starting lexState = iota
	inWord
	singleQuoting
	doubleQuoting
)

const (
	spaceRune = ' '
	singleQuote = '\''
	doubleQuote = '"'
)

func tokenize(str string) (tokens []string) {
	str = strings.TrimSpace(str)
	state := starting

	for idx, rune := range str {
		switch state {
			case starting:
				switch rune {
					case spaceRune:
						continue
					case singleQuote:
						state = singleQuoting
					case doubleQuote:
						state = doubleQuoting
					default:
						state = inWord
						if len(tokens) == 0 {
							tokens = append(tokens, "")
						}
						tokens[len(tokens) - 1] += string(rune)

				}
			case singleQuoting:
				if rune == singleQuote {
					state = inWord
					continue
				}
				if len(tokens) == 0 {
					tokens = append(tokens, "")
				}
				tokens[len(tokens) - 1] += string(rune)
			case doubleQuoting:
				if rune == doubleQuote {
					state = inWord
					continue
				}
				if len(tokens) == 0 {
					tokens = append(tokens, "")
				}
				tokens[len(tokens) - 1] += string(rune)
			case inWord:
				if rune == singleQuote {
					state = singleQuoting
					continue
				} else if rune == doubleQuote {
					state = doubleQuoting
					continue
				} else if rune == spaceRune {
				 	if str[idx - 1] == spaceRune {
						// skip
					} else {
						tokens = append(tokens, "")
					}
					continue
				}
				if len(tokens) == 0 {
					tokens = append(tokens, "")
				}
				tokens[len(tokens) - 1] += string(rune)
			default:
				fmt.Println("Unexpected state: ", state)
		}
	}
	return tokens
}
