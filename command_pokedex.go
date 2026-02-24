package main

import (
	"fmt"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

func commandPokedex(arg string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	fmt.Println("Your Pokedex:")
	for k := range *pokedex {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}
