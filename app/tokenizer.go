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
	escaping
)

const (
	spaceRune = ' '
	singleQuote = '\''
	doubleQuote = '"'
	escapeRune = '\\'
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
					case escapeRune:
						state = escaping
					default:
						state = normal
						builder.WriteRune(rune)
				}
			case singleQuoting:
				if rune == singleQuote {
					state = normal
				} else {
					builder.WriteRune(rune)
				}

			case doubleQuoting:
				if rune == doubleQuote {
					state = normal
				} else {
					builder.WriteRune(rune)
				}

			case normal:
				switch rune {
				 	case singleQuote:
							state = singleQuoting
					case doubleQuote:
						state = doubleQuoting
					case spaceRune:
					 	if str[idx - 1] == spaceRune {
							// skip
						} else {
							tokens = append(tokens, builder.String())
							builder.Reset()
						}
					case escapeRune:
						state = escaping
					default:
						builder.WriteRune(rune)
				}

			case escaping:
				builder.WriteRune(rune)
				state = normal
			default:
				fmt.Println("Unexpected state: ", state)
		}
	}
	tokens = append(tokens, builder.String())
	return tokens
}
