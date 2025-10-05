package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	repl()
}

type cliCommand struct {
	name string
	desc string
	callback func() error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand {
		"exit": {
			name: "exit",
			desc: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name: "help",
			desc: "Displays a help message",
			callback: commandHelp,
		},
	}
	return commands
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		fullCommand := cleanInput(scanner.Text())
		command, ok := getCommands()[fullCommand[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			command.callback()
		}


	}
}
