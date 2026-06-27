package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func commandMapBack(c *Config, cache *pokecache.Cache) error {
	if c.Previous == "" {
		fmt.Println("You are alredy on the first page")
		return nil
	} else {
		url := c.Previous
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
}