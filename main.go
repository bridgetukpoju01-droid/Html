package main

import (
	"fmt"
	"os"
)

func main() {
	banner, err := LoadBanner(os.Args[1])
	if err != nil {
		return
	}
	fmt.Print(GenerateArt("hi", banner))
}
