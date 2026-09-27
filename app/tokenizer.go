package main

import (
	"fmt"
	"strings"
)


type lexState int

const (
	starting lexState = iota
	normal
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
	var builder strings.Builder

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
						state = normal
						builder.WriteRune(rune)
				}
			case singleQuoting:
				if rune == singleQuote {
					state = normal
					continue
				}
				builder.WriteRune(rune)

			case doubleQuoting:
				if rune == doubleQuote {
					state = normal
					continue
				}
				builder.WriteRune(rune)

			case normal:
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
						tokens = append(tokens, builder.String())
						builder.Reset()
					}
					continue
				}
				builder.WriteRune(rune)
			default:
				fmt.Println("Unexpected state: ", state)
		}
	}
	tokens = append(tokens, builder.String())
	return tokens
}
