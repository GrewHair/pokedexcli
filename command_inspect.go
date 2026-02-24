package main

import (
	"fmt"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

func commandInspect(pokemon string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	data, found := (*pokedex)[pokemon]
	if !found {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon)
	fmt.Printf("Height: %v\n", data.Height)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Printf("Types:\n")
	for i := range data.Types {
		fmt.Printf("  - %v\n", data.Types[i].Type.Name)
	}
	return nil
}
