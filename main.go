package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// global variable for all cli commands
var cliCommands map[string]cliCommand

func cleanInput(text string) []string {
	lowerCaseText := strings.ToLower(text)
	return strings.Fields(lowerCaseText)
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Your device supports the following functionality:\n")
	fmt.Print("Usage:\t<command>: <description>\n\n")
	usageString := ""
	for key := range cliCommands {
		command := cliCommands[key]
		name := command.name
		description := command.description
		usageString += fmt.Sprintf("\t%s: %s\n", name, description)
	}
	fmt.Print(usageString)
	return nil
}

func main() {
	cliCommands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	replScanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		replScanner.Scan()
		words := cleanInput(replScanner.Text())
		command := words[0]
		c, exists := cliCommands[command]
		if exists {
			c.callback()
		} else {
			fmt.Printf("Unknown command: <%s>", command)
		}
	}
}
