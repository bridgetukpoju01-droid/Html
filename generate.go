package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	var result strings.Builder
	if input == "" {
		return ""
	}
	if input == "\\n" {
		return "\n"
	}
	text := SplitInput(input)
	for i, words := range text {
		if words == "" {
			if i != len(words)-1 {
				result.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, ch := range words {
				result.WriteString(banner[ch][row])
			}
			result.WriteString("\n")
		}
	}
	return result.String()
}
