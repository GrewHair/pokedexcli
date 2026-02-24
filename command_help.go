package main

import (
	"fmt"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

func commandHelp(arg string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
