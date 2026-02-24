package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GrewHair/pokedexcli/internal/pokecache"
)

type locationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(area string, cfg *config, cache *pokecache.Cache, pokedex *map[string]Pokemon) error {
	if area == "" {
		fmt.Println("specify an area!")
		return nil
	}

	var responseStruct locationArea
	var responseBytes []byte
	url := "https://pokeapi.co/api/v2/location-area/" + area

	cachedResponseBytes, hit := cache.Get(url)
	if hit {
		fmt.Println("=== Cached result ===")
		responseBytes = cachedResponseBytes
	} else {
		responseStream, err := http.Get(url)
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

		cache.Add(url, responseBytes)
	}

	err := json.Unmarshal(responseBytes, &responseStruct)
	if err != nil {
		fmt.Println("[ERROR]", err)
		return err
	}

	for i := range responseStruct.PokemonEncounters {
		fmt.Printf("- %s\n", responseStruct.PokemonEncounters[i].Pokemon.Name)
	}

	return nil
}
