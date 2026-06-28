package main

import (
	"fmt"
	"os"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func commandExit(c *Config, cache *pokecache.Cache, loc ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}