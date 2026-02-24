package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string, *config, *pokecache.Cache, *map[string]Pokemon) error
}

type config struct {
	nextURL     string
	previousURL string
}

func repl() {
	cfg := &config{
		nextURL:     "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		previousURL: "",
	}

	pokedex := map[string]Pokemon{}

	cache := pokecache.NewCache(30 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex 󰐝 ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandArg := ""

		if len(words) > 1 {
			commandArg = words[1]
		}

		if command, ok := getCommands()[commandName]; ok {
			command.callback(commandArg, cfg, cache, &pokedex)
		} else {
			fmt.Printf("No such command: %s\n", commandName)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Repeat for next 20",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of 20 location areas in the Pokemon world (backwards)",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores an area for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words := strings.Fields(lowercase)
	return words
}
