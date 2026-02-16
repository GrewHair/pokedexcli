package main

import (
	"net/http"
	"encoding/json"
	//"io"
	"fmt"

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

func commandMap(cfg *config, cache *pokecache.Cache) error {

	var data locationAreas

	fmt.Println("ping")
	entry, hit := cache.Get(cfg.nextUrl)
	fmt.Println("pong")
	if hit {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return err
		}
	} else {
	  res, err := http.Get(cfg.nextUrl)
	  if err != nil {
	  	return err
	  }
	  defer res.Body.Close()

	  decoder := json.NewDecoder(res.Body)

	  err = decoder.Decode(&data)
	  if err != nil {
	  	return err
	  }
  }

	cfg.nextUrl     = data.Next
	cfg.previousUrl = data.Previous

	//fmt.Printf("%#v\n", data.Results[0])
	for i := range data.Results {
	  fmt.Printf("%s\n", data.Results[i].Name)
	}
	return nil
}

func commandMapb(cfg *config, cache *pokecache.Cache) error {
	if cfg.previousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	var data locationAreas

	entry, hit := cache.Get(cfg.previousUrl)

	if hit {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return err
		}
	} else {
	  res, err := http.Get(cfg.previousUrl)
	  if err != nil {
	  	return err
	  }
	  defer res.Body.Close()

	  decoder := json.NewDecoder(res.Body)

	  err = decoder.Decode(&data)
	  if err != nil {
	  	return err
	  }
  }

	cfg.nextUrl     = data.Next
	cfg.previousUrl = data.Previous

	//fmt.Printf("%#v\n", data.Results[0])
	for i := range data.Results {
	  fmt.Printf("%s\n", data.Results[i].Name)
	}

	return nil
}
