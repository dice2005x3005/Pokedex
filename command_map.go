package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func commandMap(c *Config) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=0&limit=20")
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("The request to the api failed")
	}
	var location Location
	decoder := json.NewDecoder(res.Body)
	if er := decoder.Decode(&location); er != nil {
		fmt.Errorf("The encoding failed")
		return er
	}
	res.Body.Close()
	for i := 0; i < len(location.Results); i++ {
		fmt.Println(location.Results[i].Name)
	}
	return nil
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