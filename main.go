package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	return strings.Fields(lowerCaseText)
}

func main() {
	fmt.Println("Hello, World!")
}
