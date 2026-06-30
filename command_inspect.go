package main

import (
	"fmt"
	"github.com/dice2005x3005/Pokedex/internal/pokecache"
)

func commandInspect(c *Config, cache *pokecache.Cache, u *User, loc ...string) error {
	info, exist := u.capturas[loc[0]]
	if exist {
		fmt.Println("Name:", info.Name)
		fmt.Println("Height:", info.Height)
		fmt.Println("Weight:", info.Weight)
		fmt.Println("Stats:")
		for _, stat := range info.Stats {
			fmt.Printf("  - %s\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range info.Types {
			fmt.Printf("  - %s\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}