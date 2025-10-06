package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/b6sh/pokedexcli/internal/pokeapi"
)

func main() {
	mainClient := pokeapi.CreateClient(time.Minute*3)
	cfg := &config{
		deck: map[string]pokeapi.PokemonStats{},
		apiClient: mainClient,
	}
	repl(cfg)
}

type config struct {
	apiClient 	pokeapi.Client
	nextURL   	*string
	prevURL   	*string
	deck		map[string]pokeapi.PokemonStats
}

func repl(cfg *config) {
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
			err := command.callback(cfg, fullCommand)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

type cliCommand struct {
	name     string
	desc     string
	callback func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:     "exit",
			desc:     "Exit the Pokedex",
			callback: commandExit,
		},
		"help": {
			name:     "help",
			desc:     "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name:     "map",
			desc:     "Next locations",
			callback: commandMapf,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Prev Locations",
			callback: commandMapb,
		},
		"explore": {
			name:     "explore",
			desc:     "Explore an Area",
			callback: commandExplore,
		},
		"catch": {
			name:     "catch",
			desc:     "Catch a Pokemon!",
			callback: commandCatch,			
		},
		"inspect": {
			name:     "inspect",
			desc:     "Inspect a Pokemon",
			callback: commandInspect,			
		},
	}
	return commands
}
