package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

)

type cliCommand struct {
	name     string
	desc     string
	callback func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			desc: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			desc: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
	
func main() {


	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		inputs := cleanInput(text)

		val, ok := getCommands()[inputs[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			val.callback()
		}
	}

}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")

	for _, command := range getCommands() {
		fmt.Println(command.name + ": " + command.desc)
	}
	return nil
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	split := strings.Fields(lower)
	return split
}
