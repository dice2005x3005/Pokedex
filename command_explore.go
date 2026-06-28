package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
	"io"
)

func commandExplore(c *Config, cache *pokecache.Cache, u *User, loc ...string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", loc[0])
	if values, ok := cache.Get(url); ok != true {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		var poke PokemonEncounter
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		cache.Add(url, data)
		if err := json.Unmarshal(data, &poke); err != nil {
			return err
		}
		for i := 0; i < len(poke.PokemonEncounters); i++ {
			fmt.Println("-", poke.PokemonEncounters[i].Pokemon.Name)
		}
		return nil
	} else {
		var poke PokemonEncounter
		if err := json.Unmarshal(values, &poke); err != nil {
			return err
		}
		for i := 0; i < len(poke.PokemonEncounters); i++ {
			fmt.Println("-", poke.PokemonEncounters[i].Pokemon.Name)
		}
		return nil
	}
}

type PokemonEncounter struct {
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
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
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