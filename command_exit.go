package main

import (
	"fmt"
	"os"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

func commandExit(arg string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
