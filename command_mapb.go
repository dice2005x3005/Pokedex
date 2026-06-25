package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func commandMapBack(c *Config) error {
	if c.Previous == "" {
		fmt.Println("You are alredy on the first page")
		return nil
	} else {
		url := c.Previous
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
		c.Next = location.Next
		c.Previous = location.Previous
		return nil
	}
}