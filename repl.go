package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *pokecache.Cache) error
}

type config struct  {
	nextUrl     string
	previousUrl string
}

func repl() {
	cfg := &config{
		nextUrl: "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		previousUrl: "",
	}

	//cache := &pokecache.NewCache(5 * time.Second)
	cache := pokecache.NewCache(5 * time.Second)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex 󰐝 ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, ok := getCommands()[commandName]; ok {
			command.callback(cfg, cache)
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
	  "exit": {
	  	name:        "exit",
	  	description: "Exit the Pokedex",
	  	callback:    commandExit,
	  },
	}
}

func cleanInput(text string) []string {
	lowercase := strings.ToLower(text)
	words  := strings.Fields(lowercase)
	return words
}

