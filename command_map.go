package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func commandMap(c *Config, cache *pokecache.Cache, u *User, loc ...string) error {
	var url string
	if c.Next == "" {
		url = fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20")
	} else {
		url = c.Next
	}
	if values, ok := cache.Get(url); ok != true {
		res, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("The request to the api failed")
		}
		defer res.Body.Close()
		var location Location
		data, err := io.ReadAll(res.Body) 
		if err != nil {
			return err
		}
		cache.Add(url, data)
		if er := json.Unmarshal(data, &location); er != nil {
			return er
		}
		for i := 0; i < len(location.Results); i++ {
			fmt.Println(location.Results[i].Name)
		}
		c.Next = location.Next
		c.Previous = location.Previous
		return nil
	} else {
		var location Location
		if er := json.Unmarshal(values, &location); er != nil {
			return er
		}
		for i := 0; i < len(location.Results); i++ {
			fmt.Println(location.Results[i].Name)
		}
		return nil
	}

}

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Config struct {
	Next string
	Previous string
}