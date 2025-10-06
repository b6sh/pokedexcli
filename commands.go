package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.desc)
	}
	return nil
}

func commandMapf(cfg *config, args []string) error {
	resp, err := cfg.apiClient.ListLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.prevURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.apiClient.ListLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = resp.Next
	cfg.prevURL = resp.Previous

	for _, location := range resp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 1 {
		return errors.New("please choose an area to explore 'explore <Area Name>'")
	}
	fmt.Println("Exploring " + args[1] + "...")

	resp, err := cfg.apiClient.Explore(args[1])
	if err != nil {
		return err
	}

	for _, encounter := range resp.PokemonEncounters {
		fmt.Println(" - " + encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) == 1 {
		return errors.New("please choose a Pokemon to catch 'catch <Pokemon Name>'")
	}

	fmt.Println("Throwing a Pokeball at " + args[1] + "...")

	resp, err := cfg.apiClient.Catch(args[1])
	if err != nil {
		return err
	}
	baseXP := resp.BaseExperience

	randomNum := rand.Intn(1000)

	if randomNum < baseXP {
		fmt.Println(args[1] + " escaped!")
	} else {
		fmt.Println(args[1] + " was caught!")
	}

	cfg.deck[args[1]] = resp
	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 1 {
		return errors.New("please specify a Pokemon name 'inspect <Pokemon Name>'")
	}

	pokemon, ok := cfg.deck[args[1]]
	if !ok {
		return errors.New("you have not cought that Pokemon! try your luck with the catch command")
	}

	fmt.Println("Name: " + pokemon.Name)
	fmt.Printf("Power Level: %v\n", pokemon.BaseExperience)

	return nil
}

func commandDeck(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for key := range cfg.deck {
		fmt.Println(" - " + key)
	}

	return nil
}
