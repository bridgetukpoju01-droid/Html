package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	font := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return font, errors.New("empty file")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 856 {
		return font, errors.New("invalid file content")
	}
	for ch := ' '; ch <= '~'; ch++ {
		start := (int(ch) - 32) * 9
		font[ch] = lines[start+1 : start+9]
	}
	return font, nil
}
