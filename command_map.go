package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(arg string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	var responseStruct locationAreas
	var responseBytes []byte

	cachedResponseBytes, hit := cache.Get(cfg.nextURL)
	if hit {
		fmt.Println("=== Cached result ===")
		responseBytes = cachedResponseBytes
	} else {
		responseStream, err := http.Get(cfg.nextURL)
		if err != nil {
			fmt.Println("[ERROR]", err)
			return err
		}
		defer responseStream.Body.Close()

		responseBytes, err = io.ReadAll(responseStream.Body)
		if err != nil {
			fmt.Println("[ERROR]", err)
			return err
		}

		cache.Add(cfg.nextURL, responseBytes)
	}

	err := json.Unmarshal(responseBytes, &responseStruct)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return err
	}

	cfg.nextURL = responseStruct.Next
	cfg.previousURL = responseStruct.Previous

	// fmt.Printf("%#v\n", data.Results[0])
	for i := range responseStruct.Results {
		fmt.Printf("%s\n", responseStruct.Results[i].Name)
	}

	return nil
}

func commandMapb(arg string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	if cfg.previousURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var responseStruct locationAreas
	var responseBytes []byte

	cachedResponseBytes, hit := cache.Get(cfg.previousURL)
	if hit {
		fmt.Println("=== Cached result ===")
		responseBytes = cachedResponseBytes
	} else {
		responseStream, err := http.Get(cfg.previousURL)
		if err != nil {
			fmt.Println("[ERROR]", err)
			return err
		}
		defer responseStream.Body.Close()

		responseBytes, err := io.ReadAll(responseStream.Body)
		if err != nil {
			fmt.Println("[ERROR]", err)
			return err
		}

		cache.Add(cfg.nextURL, responseBytes)
	}

	err := json.Unmarshal(responseBytes, &responseStruct)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return err
	}

	cfg.nextURL = responseStruct.Next
	cfg.previousURL = responseStruct.Previous

	// fmt.Printf("%#v\n", data.Results[0])
	for i := range responseStruct.Results {
		fmt.Printf("%s\n", responseStruct.Results[i].Name)
	}

	return nil
}
