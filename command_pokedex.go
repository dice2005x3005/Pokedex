package main

import (
	"fmt"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func commandPokedex(c *Config, cache *pokecache.Cache, u *User, loc ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range u.capturas {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}