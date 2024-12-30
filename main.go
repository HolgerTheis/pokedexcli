package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	return strings.Fields(lowerCaseText)
}

func main() {
	replScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokédex > ")
		replScanner.Scan()
		words := cleanInput(replScanner.Text())
		fmt.Println(fmt.Sprintf("Your command was: %s", words[0]))
	}
}
