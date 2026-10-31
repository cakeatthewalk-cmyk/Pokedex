package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/cakeatthewalk-cmyk/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	var cleaned []string
	trimmedText := strings.ToLower(strings.TrimSpace(text))

	cleaned = strings.Fields(trimmedText)

	return cleaned
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:`)
	for _, value := range getAvailableCommands() {
		fmt.Printf("%s: %s\n", value.name, value.description)
	}
	return nil
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMapB(cfg *config, args []string) error {
	url := ""
	if cfg.previous != "" {
		url = cfg.previous
	}
	if url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	lo, err := cfg.client.ListLocations(url)
	if err != nil {
		return err
	}
	for _, lr := range lo.Results {
		fmt.Println(lr.Name)
	}
	cfg.next = lo.Next
	cfg.previous = lo.Previous
	return nil
}

func commandMap(cfg *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != "" {
		url = cfg.next
	}
	lo, err := cfg.client.ListLocations(url)

	if err != nil {
		return err
	}
	for _, lr := range lo.Results {
		fmt.Println(lr.Name)
	}
	cfg.next = lo.Next
	cfg.previous = lo.Previous
	return nil
}

func commandExplore(cfg *config, args []string) error {
	locationToExplore := args[0]

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", locationToExplore)

	eas, err := cfg.client.ListAreaPokemon(locationToExplore)
	if err != nil {
		return err
	}

	for _, ea := range eas.PokemonEncounters {
		fmt.Printf(" - %s\n", ea.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	pokemonName := args[0]
	fmt.Printf("Throwxing a Pokeball at %s...\n", pokemonName)
	DiceThrow := rand.Intn(100)

	if DiceThrow >= 50 {
		fmt.Printf("%s was caught!", pokemonName)
	} else {
		fmt.Printf("%s escaped!", pokemonName)
	}

	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getAvailableCommands() map[string]cliCommand {
	var availableCommands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists locations in Pokemon",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes back one page in locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "explores an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "catches a pokemon",
			callback:    commandCatch,
		},
	}
	return availableCommands
}

type config struct {
	commands map[string]cliCommand
	next     string
	previous string
	client   pokeapi.Client
}
