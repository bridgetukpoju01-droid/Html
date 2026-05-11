package main

import "fmt"

func ValidateInput(input string) (rune, error) {
	for _, i := range input {
		if i < ' ' || i > '~' {
			return i, fmt.Errorf("unsupported character: %v", i)
		}
	}
	return 0, nil
}
